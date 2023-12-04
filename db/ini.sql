
DROP DATABASE IF EXISTS BDSpotPer;

DROP TABLESPACE IF EXISTS spot_primary;
DROP TABLESPACE IF EXISTS spot_secondary;
DROP TABLESPACE IF EXISTS spot_tertiary;

CREATE TABLESPACE spot_primary
OWNER postgres
LOCATION '/pgdata/BDSpotPer/primary';

CREATE TABLESPACE spot_secondary
OWNER postgres
LOCATION '/pgdata/BDSpotPer/secondary';

CREATE TABLESPACE spot_tertiary
OWNER postgres
LOCATION '/pgdata/BDSpotPer/tertiary';

CREATE DATABASE BDSpotPer
WITH
    OWNER postgres
    TEMPLATE template0
    ENCODING 'UTF8'
    LC_COLLATE 'pt_BR.UTF-8'
    LC_CTYPE 'pt_BR.UTF-8'
    TABLESPACE = spot_primary
    ALLOW_CONNECTIONS true
    CONNECTION LIMIT = -1
    IS_TEMPLATE = false;

\connect bdspotper

CREATE TABLE gravadora (
    cod_grav   SERIAL PRIMARY KEY,
    nome       VARCHAR(20) NOT NULL,
    cidade     VARCHAR(20) NOT NULL,
    pais       VARCHAR(20) NOT NULL,
    end_homep  TEXT DEFAULT ''
) TABLESPACE spot_tertiary;

CREATE TABLE telefone_gravadora (
    cod_grav    INT,
    num         CHAR(15),
    PRIMARY KEY (cod_grav, num),
    CONSTRAINT fk_cod_grav
        FOREIGN KEY (cod_grav)
            REFERENCES gravadora(cod_grav)
            ON DELETE CASCADE
) TABLESPACE spot_tertiary;

CREATE TYPE meio_fisico AS ENUM ('CD', 'VINIL', 'DOWNLOAD');

Create Table album (
    cod_alb     SERIAL,
    cod_meio    INT,
    cod_grav    INT,
    nome        VARCHAR(20) NOT NULL,
    descricao   TEXT DEFAULT '',
    data_grav   DATE NOT NULL CHECK (data_grav > '2000-01-01'),
    pr_compra   DECIMAL (10,2) NOT NULL,
    pr_venda    DECIMAL (10,2),
    meio        meio_fisico NOT NULL,
    PRIMARY KEY (cod_alb, cod_meio),
    CONSTRAINT fk_cod_grav
        FOREIGN KEY (cod_grav)
            REFERENCES gravadora(cod_grav)
            ON DELETE SET NULL
) TABLESPACE spot_tertiary;

CREATE TYPE forma_musical AS ENUM (
    'Sinfonia', 'Sonata', 'Ópera', 'Concerto', 'Suite', 'Missa', 'Música de câmara',
    'Canção', 'Poema sinfônico', 'Rapsódia', 'Fuga', 'Réquiem'
);

CREATE TABLE composicao (
    cod_composicao   SERIAL PRIMARY KEY,
    descricao        TEXT DEFAULT '',
    tipo             forma_musical NOT NULL
) TABLESPACE spot_tertiary;

CREATE TYPE tipo_interprete AS ENUM (
    'Orquestra', 'Trio', 'Quarteto', 'Ensemble', 'Soprano', 'Tenor', 'Ópera'
);

CREATE TABLE interprete (
    cod_inter   SERIAL PRIMARY KEY,
    nome        VARCHAR(20) NOT NULL,
    tipo        tipo_interprete NOT NULL
) TABLESPACE spot_tertiary;

CREATE TYPE periodo_historico as ENUM
('idade média', 'renascença', 'barroco', 'clássico', 'romântico', 'moderno');

CREATE TABLE periodo_musical (
    cod_pm      SERIAL PRIMARY KEY,
    periodo     periodo_historico,
    int_tempo   DATERANGE NOT NULL
) TABLESPACE spot_tertiary;

CREATE TABLE compositor (
    cod_compositor  SERIAL PRIMARY KEY,

    cod_pm          INT NOT NULL,
    
    nome            VARCHAR(20) NOT NULL,
    dt_nasc         DATE NOT NULL,
    dt_morte        DATE,
    CONSTRAINT fk_cod_pm
        FOREIGN KEY (cod_pm)
            REFERENCES periodo_musical(cod_pm)
                ON DELETE NO ACTION
) TABLESPACE spot_tertiary;


CREATE TABLE playlist (
    cod_play    SERIAL PRIMARY KEY,
    nome        VARCHAR(20) NOT NULL,
    tempo_tot   TIME DEFAULT '00:00:00',
    data_criac  DATE DEFAULT CURRENT_DATE
) TABLESPACE spot_secondary;

CREATE TYPE gravacao AS ENUM ( 'ADD', 'DDD' );

CREATE TABLE faixa (
    cod_faixa   SERIAL,
    
    cod_alb     INT NOT NULL,
    cod_meio    INT NOT NULL,
    cod_composicao  INT NOT NULL,

    numero      INT UNIQUE NOT NULL,
    descricao   TEXT DEFAULT '',
    tempo_exec  TIME NOT NULL,
    tipo_grav   gravacao,
    
    CONSTRAINT fk_cod_composicao
        FOREIGN KEY (cod_composicao)
            REFERENCES composicao(cod_composicao)
            ON DELETE CASCADE,
    CONSTRAINT pk_album_faixa
        FOREIGN KEY (cod_alb, cod_meio)
            REFERENCES album(cod_alb, cod_meio)
            ON DELETE CASCADE,
    PRIMARY KEY (cod_alb, cod_meio, cod_faixa)
) TABLESPACE spot_secondary;

CREATE TABLE faixa_interprete (
    cod_faixa   INT NOT NULL,
    cod_alb     INT NOT NULL,
    cod_meio    INT NOT NULL,

    cod_inter   INT,
    CONSTRAINT fk_cod_faixa
        FOREIGN KEY (cod_faixa, cod_alb, cod_meio)
            REFERENCES faixa(cod_faixa, cod_alb, cod_meio)
            ON DELETE CASCADE,
    CONSTRAINT fk_cod_inter
        FOREIGN KEY (cod_inter)
            REFERENCES interprete(cod_inter),
    PRIMARY KEY (cod_faixa, cod_alb, cod_meio, cod_inter)
) TABLESPACE spot_tertiary;


CREATE TABLE faixa_compositor (
    cod_faixa       INT NOT NULL,
    cod_alb         INT NOT NULL,
    cod_meio        INT NOT NULL,

    cod_compositor  INT,
    CONSTRAINT fk_cod_faixa
        FOREIGN KEY (cod_faixa, cod_alb, cod_meio)
            REFERENCES faixa(cod_faixa, cod_alb, cod_meio)
            ON DELETE CASCADE,
    CONSTRAINT fk_cod_compositor
        FOREIGN KEY (cod_compositor)
            REFERENCES compositor(cod_compositor),
    PRIMARY KEY (cod_faixa, cod_alb, cod_meio, cod_compositor)
) TABLESPACE spot_tertiary;

CREATE TABLE faixa_playlist (
    cod_faixa       INT NOT NULL,
    cod_alb         INT NOT NULL,
    cod_meio        INT NOT NULL,

    cod_play        INT NOT NULL,
    dt_ult_repr     DATE,
    qtd_repr        INT DEFAULT 0,
    CONSTRAINT fk_cod_faixa
        FOREIGN KEY (cod_faixa, cod_alb, cod_meio)
            REFERENCES faixa(cod_faixa, cod_alb, cod_meio)
            ON DELETE CASCADE,
    CONSTRAINT fk_cod_play
        FOREIGN KEY (cod_play)
            REFERENCES playlist(cod_play),
    PRIMARY KEY (cod_faixa, cod_alb, cod_meio, cod_play)
) TABLESPACE spot_secondary;

