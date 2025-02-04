-- Buat ENUM type untuk jenis jaringan blockchain jika belum ada
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'network_type') THEN
        CREATE TYPE network_type AS ENUM ('MAINNET', 'TESTNET', 'PRIVATE');
    END IF;
END $$ LANGUAGE plpgsql;

-- Buat tabel untuk menyimpan informasi jaringan blockchain
CREATE TABLE IF NOT EXISTS networks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,  -- Nama jaringan (Ethereum, Polygon, BSC, dll.)
    chain_id BIGINT NOT NULL UNIQUE,    -- Chain ID jaringan (<0 untuk jaringan PRIVATE)
    rpc_url TEXT NOT NULL,              -- URL RPC jaringan (Infura, Alchemy, dll.)
    explorer_url TEXT,                   -- URL block explorer (Opsional)
    symbol VARCHAR(10) NOT NULL,         -- Simbol token utama (ETH, MATIC, BNB, dll.)
    type network_type DEFAULT 'MAINNET', -- ENUM untuk jenis jaringan
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert data jaringan blockchain utama & PRIVATE
INSERT INTO networks (name, chain_id, rpc_url, explorer_url, symbol,type)
VALUES 
    -- 🌍 PUBLIC MAINNET NETWORKS
    ('Ethereum MAINNET', 1, 'https://MAINNET.infura.io/v3/YOUR_INFURA_KEY', 'https://etherscan.io', 'ETH', 'MAINNET'),
    ('Binance Smart Chain MAINNET', 56, 'https://bsc-dataseed.binance.org', 'https://bscscan.com', 'BNB', 'MAINNET'),
    ('Polygon MAINNET', 137, 'https://polygon-rpc.com', 'https://polygonscan.com', 'MATIC', 'MAINNET'),
    ('Avalanche C-Chain', 43114, 'https://api.avax.network/ext/bc/C/rpc', 'https://snowtrace.io', 'AVAX', 'MAINNET'),
    ('Fantom Opera', 250, 'https://rpcapi.fantom.network', 'https://ftmscan.com', 'FTM', 'MAINNET'),
    ('Arbitrum One', 42161, 'https://arb1.arbitrum.io/rpc', 'https://arbiscan.io', 'ETH', 'MAINNET'),
    ('Optimism MAINNET', 10, 'https://MAINNET.optimism.io', 'https://optimistic.etherscan.io', 'ETH', 'MAINNET'),
    ('Cronos MAINNET', 25, 'https://evm.cronos.org', 'https://cronoscan.com', 'CRO','MAINNET'),
    
    -- 🔬 TESTNET NETWORKS
    ('Ethereum Goerli', 5, 'https://goerli.infura.io/v3/YOUR_INFURA_KEY', 'https://goerli.etherscan.io', 'ETH','TESTNET'),
    ('Ethereum Sepolia', 11155111, 'https://sepolia.infura.io/v3/YOUR_INFURA_KEY', 'https://sepolia.etherscan.io', 'ETH','TESTNET'),
    ('Polygon Mumbai', 80001, 'https://rpc-mumbai.maticvigil.com', 'https://mumbai.polygonscan.com', 'MATIC','TESTNET'),
    ('Avalanche Fuji', 43113, 'https://api.avax-test.network/ext/bc/C/rpc', 'https://TESTNET.snowtrace.io', 'AVAX','TESTNET'),
    ('Fantom TESTNET', 4002, 'https://rpc.TESTNET.fantom.network', 'https://TESTNET.ftmscan.com', 'FTM','TESTNET'),
    ('Arbitrum Goerli', 421613, 'https://goerli-rollup.arbitrum.io/rpc', 'https://goerli.arbiscan.io', 'ETH','TESTNET'),
    ('Optimism Goerli', 420, 'https://goerli.optimism.io', 'https://goerli-optimistic.etherscan.io', 'ETH','TESTNET'),

    -- 🔐 PRIVATE NETWORKS (Quorum, Hyperledger Fabric, dll.)
    ('Quorum PRIVATE Network', -1, 'http://127.0.0.1:8545', NULL, 'ETH', 'PRIVATE'),
    ('Hyperledger Fabric PRIVATE', -2, 'grpc://127.0.0.1:7051', NULL, 'FABRIC', 'PRIVATE'),
    ('Corda PRIVATE Network', -3, 'http://127.0.0.1:10050', NULL, 'CORDA', 'PRIVATE'),
    ('Binance Smart Chain PRIVATE', -4, 'http://127.0.0.1:8545', NULL, 'BNB', 'PRIVATE'),
    ('Polygon Edge PRIVATE', -5, 'http://127.0.0.1:10002', NULL, 'MATIC', 'PRIVATE'),
    ('Avalanche Subnet PRIVATE', -6, 'http://127.0.0.1:9650/ext/bc/C/rpc', NULL, 'AVAX', 'PRIVATE')
ON CONFLICT (chain_id) DO NOTHING;
