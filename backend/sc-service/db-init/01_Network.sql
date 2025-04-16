CREATE SCHEMA IF NOT EXISTS ref;
-- Buat ENUM type untuk jenis jaringan blockchain jika belum ada

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'network_type') THEN
        CREATE TYPE network_type AS ENUM ('mainnet', 'testnet', 'private');
    END IF;
END $$ LANGUAGE plpgsql;



CREATE TABLE IF NOT EXISTS ref.blockchain_platform (
	"id" UUID NOT NULL,
	"nm_blockchain" VARCHAR(50) NOT NULL,
	"active" BOOLEAN NULL DEFAULT 'false',
	"created_at" TIMESTAMP NULL DEFAULT 'now()',
	"updated_at" TIMESTAMP NULL DEFAULT 'now()',
	PRIMARY KEY ("id")
);

INSERT INTO ref.blockchain_platform ("id", "nm_blockchain", "active", "created_at", "updated_at") VALUES ('f45865b2-9dd9-4085-942c-89a8d1847674', 'Ethereum', 'false', '2025-04-09 16:09:11.383076', '2025-04-09 16:09:11.383076');
INSERT INTO ref.blockchain_platform ("id", "nm_blockchain", "active", "created_at", "updated_at") VALUES ('dd0614c8-84ef-4a49-90c6-91540f1ed4aa', 'Quorum', 'false', '2025-04-09 16:09:11.383076', '2025-04-09 16:09:11.383076');
INSERT INTO ref.blockchain_platform ("id", "nm_blockchain", "active", "created_at", "updated_at") VALUES ('1947d865-0a2a-4903-a3c7-42ce59cecb39', 'Hyperledger Fabric', 'false', '2025-04-09 16:09:11.383076', '2025-04-09 16:09:11.383076');


CREATE TABLE IF NOT EXISTS ref.networks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,  -- Nama jaringan (Ethereum, Polygon, BSC, dll.)
    chain_id BIGINT NOT NULL UNIQUE,    -- Chain ID jaringan (<0 untuk jaringan PRIVATE)
    rpc_url TEXT NOT NULL,              -- URL RPC jaringan (Infura, Alchemy, dll.)
    explorer_url TEXT,                   -- URL block explorer (Opsional)
    symbol VARCHAR(10) NOT NULL,         -- Simbol token utama (ETH, MATIC, BNB, dll.)
    type network_type DEFAULT 'mainnet', -- ENUM untuk jenis jaringan
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    activate BOOLEAN NULL DEFAULT false,
	 available BOOLEAN NULL DEFAULT false,
    architecture VARCHAR(100) NULL DEFAULT 'EVM'
);

-- Insert data jaringan blockchain utama & PRIVATE
INSERT INTO ref.networks (name, chain_id, rpc_url, explorer_url, symbol,TYPE, activate, available,architecture)
VALUES 
    -- 🌍 PUBLIC MAINNET NETWORKS
    ('Ethereum MAINNET', 1, 'https://MAINNET.infura.io/v3/YOUR_INFURA_KEY', 'https://etherscan.io', 'ETH', 'mainnet', false,false,'EVM'),
    ('Binance Smart Chain MAINNET', 56, 'https://bsc-dataseed.binance.org', 'https://bscscan.com', 'BNB', 'mainnet', false,false,'EVM'),
    ('Polygon MAINNET', 137, 'https://polygon-rpc.com', 'https://polygonscan.com', 'MATIC', 'mainnet', false,false,'EVM'),
    ('Avalanche C-Chain', 43114, 'https://api.avax.network/ext/bc/C/rpc', 'https://snowtrace.io', 'AVAX', 'mainnet', false,false,'EVM'),
    ('Fantom Opera', 250, 'https://rpcapi.fantom.network', 'https://ftmscan.com', 'FTM', 'mainnet', false,false,'EVM'),
    ('Arbitrum One', 42161, 'https://arb1.arbitrum.io/rpc', 'https://arbiscan.io', 'ETH', 'mainnet', false,false,'EVM'),
    ('Optimism MAINNET', 10, 'https://MAINNET.optimism.io', 'https://optimistic.etherscan.io', 'ETH', 'mainnet', false,false,'EVM'),
    ('Cronos MAINNET', 25, 'https://evm.cronos.org', 'https://cronoscan.com', 'CRO','mainnet', false,false,'EVM'),
    
    -- 🔬 TESTNET NETWORKS
    ('Ethereum Goerli', 5, 'https://goerli.infura.io/v3/YOUR_INFURA_KEY', 'https://goerli.etherscan.io', 'ETH','testnet', false,false,'EVM'),
    ('Ethereum Sepolia', 11155111, 'https://sepolia.infura.io/v3/YOUR_INFURA_KEY', 'https://sepolia.etherscan.io', 'ETH','testnet', false,false,'EVM'),
    ('Polygon Mumbai', 80001, 'https://rpc-mumbai.maticvigil.com', 'https://mumbai.polygonscan.com', 'MATIC','testnet', false,false,'EVM'),
    ('Avalanche Fuji', 43113, 'https://api.avax-test.network/ext/bc/C/rpc', 'https://TESTNET.snowtrace.io', 'AVAX','testnet', false,false,'EVM'),
    ('Fantom TESTNET', 4002, 'https://rpc.TESTNET.fantom.network', 'https://TESTNET.ftmscan.com', 'FTM','testnet', false,false,'EVM'),
    ('Arbitrum Goerli', 421613, 'https://goerli-rollup.arbitrum.io/rpc', 'https://goerli.arbiscan.io', 'ETH','testnet', false,false,'EVM'),
    ('Optimism Goerli', 420, 'https://goerli.optimism.io', 'https://goerli-optimistic.etherscan.io', 'ETH','testnet', false,false,'EVM'),

    -- 🔐 PRIVATE NETWORKS (Quorum, Hyperledger Fabric, dll.)
    ('Quorum PRIVATE Network', -1, 'http://127.0.0.1:8545', NULL, 'ETH', 'private', false,false,'EVM'),
    ('Hyperledger Fabric PRIVATE', -2, 'grpc://127.0.0.1:7051', NULL, 'FABRIC', 'private', false,false,'NONEVM'),
    ('Corda PRIVATE Network', -3, 'http://127.0.0.1:10050', NULL, 'CORDA', 'private', FALSE, FALSE,'NONEVM'),
    ('Binance Smart Chain PRIVATE', -4, 'http://127.0.0.1:8545', NULL, 'BNB', 'private', false,false,'EVM'),
    ('Polygon Edge PRIVATE', -5, 'http://127.0.0.1:10002', NULL, 'MATIC', 'private', false,false,'EVM'),
    ('Avalanche Subnet PRIVATE', -6, 'http://127.0.0.1:9650/ext/bc/C/rpc', NULL, 'AVAX', 'private', false,false,'EVM')
ON CONFLICT (chain_id) DO NOTHING;

