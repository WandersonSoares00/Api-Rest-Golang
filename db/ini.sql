
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

DROP SCHEMA IF EXISTS spot;
CREATE SCHEMA spot;
SET search_path TO spot, public;

\c bdspotper

CREATE TABLE telefone_gravadora (
    cod_fone    SMALLINT PRIMARY KEY,
    num         VARCHAR(12)
) TABLESPACE spot_tertiary;

Create Table gravadora (
    cod_grav   SMALLINT PRIMARY KEY,
    cod_fone   SMALLINT,
    nome       VARCHAR(20),
    endereco   VARCHAR(20),
    end_homep  VARCHAR(20),
    CONSTRAINT fk_cod_fone FOREIGN KEY (cod_fone)
        REFERENCES telefone_gravadora(cod_fone)
        ON DELETE SET NULL
) TABLESPACE spot_tertiary;

CREATE TABLE meio_download (
    cod_down    SMALLINT PRIMARY KEY
) TABLESPACE spot_tertiary;

Create Table Album (
    descricao   VARCHAR (20),
    cod_alb     SMALLINT PRIMARY KEY,
    cod_down    SMALLINT,
    data_grav   DATE,
    pr_compra   DECIMAL (10,2),
    pr_venda    DECIMAL (10,2),
    CONSTRAINT fk_cod_down FOREIGN KEY (cod_down)
        REFERENCES meio_download(cod_down)
        ON DELETE SET NULL
) TABLESPACE spot_tertiary;

CREATE TABLE meio_cd (
    cod_cd     SMALLINT PRIMARY KEY
) TABLESPACE spot_tertiary;

CREATE TABLE meio_vinil (
    cod_vinil  SMALLINT PRIMARY KEY
) TABLESPACE spot_tertiary;

Create table composicao (
    cod_composicao   SMALLINT PRIMARY KEY,
    descricao        TEXT,
    tipo             VARCHAR(20)
) TABLESPACE spot_tertiary;

Create table faixa (
    cod_faixa   SMALLINT PRIMARY KEY,
    cod_cd      SMALLINT,
    cod_vinil   SMALLINT,
    cod_down    SMALLINT,
    cod_composicao  SMALLINT NOT NULL,
    numero      INT,
    descricao   TEXT,
    tempo_exec  TIME,
    tipo_grav   VARCHAR(8),
    CONSTRAINT fk_cod_cd FOREIGN KEY (cod_cd)
        REFERENCES meio_cd(cod_cd)
        ON DELETE SET NULL,
    CONSTRAINT fk_cod_vinil FOREIGN KEY (cod_vinil)
        REFERENCES meio_vinil(cod_vinil)
        ON DELETE SET NULL,
    CONSTRAINT fk_cod_down FOREIGN KEY (cod_down)
        REFERENCES meio_download(cod_down)
        ON DELETE SET NULL,
    CONSTRAINT fk_cod_composicao FOREIGN KEY (cod_composicao)
        REFERENCES composicao(cod_composicao)
) TABLESPACE spot_secondary;

CREATE TABLE periodo_musical (
    cod_pm      SMALLINT PRIMARY KEY,
    descricao   TEXT,
    int_tempo INTERVAL
) TABLESPACE spot_tertiary;

CREATE TABLE compositor (
    cod_compositor  SMALLINT PRIMARY KEY,
    cod_pm          SMALLINT NOT NULL,
    dt_nasc         DATE NOT NULL,
    dt_morte        DATE,
    nome            VARCHAR(20) NOT NULL,
    CONSTRAINT fk_cod_pm FOREIGN KEY (cod_pm)
        REFERENCES periodo_musical(cod_pm)
) TABLESPACE spot_tertiary;


CREATE TABLE faixa_compositor (
    cod_faixa       SMALLINT NOT NULL,
    cod_compositor  SMALLINT NOT NULL,
    CONSTRAINT fk_cod_faixa FOREIGN KEY (cod_faixa)
        REFERENCES faixa(cod_faixa),
    CONSTRAINT fk_cod_compositor FOREIGN KEY (cod_compositor)
        REFERENCES compositor(cod_compositor)
) TABLESPACE spot_tertiary;


CREATE TABLE playlist (
    cod_play    SMALLINT PRIMARY KEY,
    nome        VARCHAR(20) NOT NULL,
    tempo_tot   TIME NOT NULL,
    data_criac  DATE
) TABLESPACE spot_secondary;

CREATE TABLE faixa_playlist (
    cod_faixa       SMALLINT NOT NULL,
    cod_play        SMALLINT NOT NULL,
    CONSTRAINT fk_cod_faixa FOREIGN KEY (cod_faixa)
        REFERENCES faixa(cod_faixa),
    CONSTRAINT fk_cod_play FOREIGN KEY (cod_play)
        REFERENCES playlist(cod_play)
--      ON DELETE SET NULL
) TABLESPACE spot_secondary;

CREATE TABLE interprete (
    cod_inter   SMALLINT PRIMARY KEY,
    nome        VARCHAR(20) NOT NULL,
    tipo        VARCHAR(20)
) TABLESPACE spot_tertiary;

CREATE TABLE faixa_interprete (
    cod_faixa       SMALLINT NOT NULL,
    cod_inter       SMALLINT,
    CONSTRAINT fk_cod_faixa FOREIGN KEY (cod_faixa)
        REFERENCES faixa(cod_faixa),
    CONSTRAINT fk_cod_inter FOREIGN KEY (cod_inter)
        REFERENCES interprete(cod_inter)
) TABLESPACE spot_tertiary;

