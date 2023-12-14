
/*
Quando o meio físico de armazenamento é CD, o tipo de gravação tem que
ser ADD ou DDD. Quando o meio físico de armazenamento é vinil ou
download, o tipo de gravação não terá valor algum.
*/
CREATE OR REPLACE FUNCTION check_meio_fisico (_cod_alb INT, _meio meio_fisico, _tipo_grav gravacao)
RETURNS BOOLEAN AS $$
DECLARE
    tipo_meio    meio_fisico;
BEGIN
    SELECT meio INTO tipo_meio FROM album
    WHERE meio = _meio AND cod_alb = _cod_alb;

    IF tipo_meio = 'CD' THEN RETURN (_tipo_grav = 'ADD' OR _tipo_grav = 'DDD');
    ELSE                RETURN _tipo_grav IS NULL;
    END IF;
END
$$ LANGUAGE PLPGSQL;

ALTER TABLE faixa DROP CONSTRAINT IF EXISTS meio_fisico_gravacao;
ALTER TABLE faixa ADD CONSTRAINT meio_fisico_gravacao CHECK (check_meio_fisico(cod_alb, meio, tipo_grav));

/*
Um álbum, com faixas de músicas do período barroco, só pode ser inserido no
banco de dados, caso o tipo de gravação seja DDD.
*/
CREATE OR REPLACE FUNCTION check_faixas_barroco (_cod_alb INT, _meio meio_fisico, _nro_faixa INT, _cod_comp INT)
RETURNS BOOLEAN AS $$
DECLARE
    is_barroco  BOOLEAN;
BEGIN
    SELECT TRUE into is_barroco FROM compositor
    JOIN periodo_musical USING(cod_pm)
    WHERE cod_compositor = _cod_comp AND periodo = 'barroco';

    IF is_barroco THEN
        RETURN EXISTS (
            SELECT tipo_grav FROM faixa
            WHERE cod_alb = _cod_alb AND meio = _meio AND nro_faixa = _nro_faixa AND tipo_grav = 'DDD'
        );
    ELSE
        RETURN TRUE;
    END IF;
END
$$ LANGUAGE PLPGSQL;

ALTER TABLE faixa_compositor DROP CONSTRAINT IF EXISTS limit_periodo_barroco;
ALTER TABLE faixa_compositor ADD CONSTRAINT limit_periodo_barroco CHECK (check_faixas_barroco(cod_alb, meio, nro_faixa, cod_compositor));

/*
Um álbum não pode ter mais que 64 faixas (músicas)
*/
CREATE OR REPLACE FUNCTION check_limit_faixas (_cod_alb INT, _meio meio_fisico)
RETURNS BOOLEAN AS $$
DECLARE
    qtd_faixas      INT;
BEGIN
    SELECT COUNT (nro_faixa) INTO qtd_faixas FROM album
    JOIN faixa USING (cod_alb, meio)
    WHERE cod_alb = _cod_alb AND meio = _meio;
    RETURN qtd_faixas < 64;
END
$$ LANGUAGE PLPGSQL;

ALTER TABLE album DROP CONSTRAINT IF EXISTS limit_64_faixas;
ALTER TABLE album ADD CONSTRAINT limit_64_faixas CHECK (check_limit_faixas(cod_alb, meio));

/*
O preço de compra de um álbum não deve ser superior a três vezes a média
do preço de compra de álbuns, com todas as faixas com tipo de gravação DDD.
*/
CREATE OR REPLACE FUNCTION check_limit_pr (_pr_compra DECIMAL)
RETURNS BOOLEAN AS $$
DECLARE
    preco DECIMAL(10, 2) = 1.00;
BEGIN
    SELECT AVG(pr_compra) INTO preco FROM album
    WHERE meio = 'CD';
    RETURN _pr_compra <= 3 * preco;
END
$$ LANGUAGE PLPGSQL;

ALTER TABLE album DROP CONSTRAINT IF EXISTS limit_pr_compra;
ALTER TABLE album ADD CONSTRAINT limit_pr_compra CHECK (check_limit_pr(pr_compra));

/*
Atualiza o tempo total de cada playlist após inserções e remoções das suas faixas
*/
CREATE OR REPLACE FUNCTION check_tempo_playlist()
   RETURNS TRIGGER
AS $$
DECLARE
    tempo_faixa TIME;
BEGIN
    IF (TG_OP = 'INSERT') THEN
        SELECT tempo_exec INTO tempo_faixa FROM faixa
        WHERE nro_faixa = NEW.nro_faixa AND cod_alb = NEW.cod_alb AND meio = NEW.meio;
        
        UPDATE playlist SET tempo_tot = tempo_tot + tempo_faixa::INTERVAL
        WHERE cod_play = NEW.cod_play;
    ELSIF (TG_OP = 'DELETE') THEN
        SELECT tempo_exec INTO tempo_faixa FROM faixa
        WHERE nro_faixa = OLD.nro_faixa AND cod_alb = OLD.cod_alb AND meio = OLD.meio;

        UPDATE playlist SET tempo_tot = tempo_tot - tempo_faixa
        WHERE cod_play = OLD.cod_play;
        
    END IF;

    RETURN NULL;
END
$$ LANGUAGE PLPGSQL;

CREATE OR REPLACE TRIGGER tempo_playlist
    AFTER INSERT OR DELETE
    ON faixa_playlist
    FOR EACH ROW
    EXECUTE PROCEDURE check_tempo_playlist();

