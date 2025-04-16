CREATE SCHEMA IF NOT EXISTS {{schema_name}};
CREATE TABLE IF NOT EXISTS {{schema_name}}.accounts ( 
    id BIGSERIAL PRIMARY KEY,
    address TEXT NOT NULL,
    username VARCHAR(100) NULL DEFAULT NULL,
    type TEXT NULL DEFAULT 'IMPORTED',
    user_id INTEGER NOT NULL,
    password VARCHAR(100) NULL DEFAULT NULL,
    keystroke_filename TEXT NULL DEFAULT NULL,
    network_id BIGINT NOT NULL,
    organization TEXT NULL DEFAULT NULL,
    is_active BOOLEAN NULL DEFAULT NULL,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    UNIQUE (address),
    CONSTRAINT fk_accounts_network FOREIGN KEY (network_id) REFERENCES ref.networks (id) 
        ON UPDATE CASCADE ON DELETE SET NULL
);

-- Membuat indeks secara terpisah
CREATE INDEX idx_accounts_network_id ON {{schema_name}}.accounts (network_id);
CREATE INDEX idx_accounts_user_id ON {{schema_name}}.accounts (user_id);

CREATE TABLE IF NOT EXISTS {{schema_name}}.contracts (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    contract_address TEXT NOT NULL,
    abi TEXT NOT NULL,
    bytecode TEXT,
    network_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    UNIQUE (contract_address),
    CONSTRAINT fk_contracts_network FOREIGN KEY (network_id) REFERENCES ref.networks (id) 
        ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_contracts_network_id ON {{schema_name}}.contracts (network_id);

CREATE TABLE IF NOT EXISTS {{schema_name}}.transactions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    tx_hash TEXT NOT NULL,
    contract_id BIGINT,
    network_id BIGINT NOT NULL,
    method VARCHAR(100),
    input_data TEXT,
    gas_used BIGINT,
    status VARCHAR(20),
    timestamp TIMESTAMPTZ,
    UNIQUE (tx_hash),
    CONSTRAINT fk_transactions_account FOREIGN KEY (account_id) REFERENCES {{schema_name}}.accounts (id) 
        ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_transactions_contract FOREIGN KEY (contract_id) REFERENCES {{schema_name}}.contracts (id) 
        ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_transactions_network FOREIGN KEY (network_id) REFERENCES ref.networks (id) 
        ON UPDATE CASCADE ON DELETE SET NULL
);
-- 
CREATE INDEX idx_transactions_account_id ON {{schema_name}}.transactions (account_id);
CREATE INDEX idx_transactions_network_id ON {{schema_name}}.transactions (network_id);
CREATE INDEX idx_transactions_contract_id ON {{schema_name}}.transactions (contract_id);


-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE TABLE {{schema_name}}.blockchain_platform (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     nm_blockchain VARCHAR(50) NOT NULL UNIQUE,
--     active BOOLEAN DEFAULT FALSE,
--     created_at TIMESTAMP DEFAULT NOW(),
--     updated_at TIMESTAMP DEFAULT NOW()
-- );

-- CREATE FUNCTION {{schema_name}}.update_timestamp()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.updated_at = NOW();
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER trg_update_timestamp
-- BEFORE UPDATE ON {{schema_name}}.blockchain_platform
-- FOR EACH ROW
-- EXECUTE FUNCTION {{schema_name}}.update_timestamp();

-- CREATE OR REPLACE FUNCTION {{schema_name}}.enforce_single_active_blockchain()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     -- Set semua baris lain menjadi false sebelum menyimpan yang baru sebagai true
--     IF NEW.active AND (OLD.active IS DISTINCT FROM TRUE) THEN
--         -- Kunci tabel untuk menghindari race condition
--         LOCK TABLE {{schema_name}}.blockchain_platform IN EXCLUSIVE MODE;
--          -- Set semua baris lain menjadi false sebelum menyimpan yang baru sebagai true
--         UPDATE {{schema_name}}.blockchain_platform 
--         SET active = FALSE 
--         WHERE id <> NEW.id;
--     END IF;
    
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER trigger_single_active_blockchain
-- BEFORE INSERT OR UPDATE ON {{schema_name}}.blockchain_platform
-- FOR EACH ROW
-- EXECUTE FUNCTION {{schema_name}}.enforce_single_active_blockchain();

-- INSERT INTO {{schema_name}}.blockchain_platform (nm_blockchain, active) VALUES
-- ('Ethereum', FALSE),
-- ('Quorum', FALSE),
-- ('Hyperledger Fabric', FALSE);
