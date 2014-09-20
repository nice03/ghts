DROP DATABASE IF EXISTS pts;
CREATE DATABASE pts DEFAULT CHARACTER SET = utf8mb4;

GRANT ALL ON pts.* TO 'pts'@'localhost' IDENTIFIED BY 'FqQrDsEm9f2Vw6pR' WITH GRANT OPTION;
GRANT ALL ON pts.* TO 'pts'@'127.0.0.1' IDENTIFIED BY 'FqQrDsEm9f2Vw6pR' WITH GRANT OPTION;

USE pts;

CREATE TABLE stock_info (
	id	INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	issuer  VARCHAR(150),
	name1	VARCHAR(150),
	name2	VARCHAR(150),
	code1	VARCHAR(40),
	code2	VARCHAR(40),
	market_status VARCHAR(250),
    UNIQUE INDEX idx_code1 (code1),
    INDEX idx_code2 (code2),
    INDEX idx_name1 (name1)
) ENGINE=InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE stock_daily_price (
	id	BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	stock_info_id	INT UNSIGNED,
	priced_on	DATE DEFAULT 0 NOT NULL,
	open		DECIMAL(65,5) NOT NULL DEFAULT '0',
	high		DECIMAL(65,5) NOT NULL DEFAULT '0',
	low			DECIMAL(65,5) NOT NULL DEFAULT '0',
	close		DECIMAL(65,5) NOT NULL DEFAULT '0',
    adj_coeff   DOUBLE NOT NULL DEFAULT '0',
    adj_open    DECIMAL(65,5) NOT NULL DEFAULT '0',
    adj_high	DECIMAL(65,5) NOT NULL DEFAULT '0',
	adj_low		DECIMAL(65,5) NOT NULL DEFAULT '0',
    adj_close	DECIMAL(65,5) NOT NULL DEFAULT '0',
    volumn      DOUBLE NOT NULL DEFAULT '0',
    previous_close  DECIMAL(65,5) NOT NULL DEFAULT '0',
	FOREIGN KEY (stock_info_id) REFERENCES stock_info (id),	
    INDEX idx_daily_price (stock_info_id, priced_on)	
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE error_obtaining_stock_daily_price (
    id	BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	stock_info_id	INT UNSIGNED,
	queried_on	DATE,
	FOREIGN KEY (stock_info_id) REFERENCES stock_info (id),
    INDEX idx_queried_on (queried_on)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;