
/*
função que tem como parâmetro de entrada o nome (ou parte do)
nome do compositor e o parâmetro de saída todos os álbuns com obras
compostas pelo compositor.
*/

CREATE OR REPLACE FUNCTION albuns_compositor (_nome_comp VARCHAR)
RETURNS TABLE(album VARCHAR)
AS $$
BEGIN
    RETURN query
    SELECT DISTINCT a.nome "album" FROM album a
    JOIN faixa              USING (cod_alb, cod_meio)
    JOIN faixa_compositor   USING (cod_faixa, cod_alb, cod_meio)
    JOIN compositor c       USING (cod_compositor)
    WHERE c.nome LIKE _nome_comp;
END
$$ LANGUAGE PLPGSQL;

/*
Listar os álbuns com preço de compra maior que a média de preços de
compra de todos os álbuns.
*/
CREATE OR REPLACE FUNCTION albuns_pr_acima_media ()
RETURNS TABLE(album VARCHAR)
AS $$
BEGIN
    RETURN query
    SELECT nome FROM album
        WHERE pr_compra >= all (
            SELECT avg(pr_compra) FROM album
    );
END
$$ LANGUAGE PLPGSQL;


/*
Listar nome da gravadora com maior número de playlists que possuem
pelo uma faixa composta pelo compositor Dvorack
*/
CREATE OR REPLACE FUNCTION gravadora_com_mais_playlists (_nome_comp VARCHAR DEFAULT '_vorack')
RETURNS TABLE(nome_gravadora VARCHAR)
AS $$
BEGIN
    RETURN query
    SELECT nome FROM (
        SELECT nome, MAX(qtd) FROM (
            SELECT g.nome, COUNT(distinct cod_play) qtd
            FROM gravadora g
            JOIN album          USING(cod_grav)
            JOIN faixa f        USING(cod_alb, cod_meio)
            JOIN faixa_playlist USING(cod_faixa, cod_alb, cod_meio)
            JOIN playlist       USING(cod_play)
            WHERE exists (
                SELECT 1 FROM compositor c
                JOIN faixa_compositor fc ON f.cod_faixa = fc.cod_faixa
                                         AND f.cod_alb  = fc.cod_alb
                                         AND f.cod_meio = fc.cod_meio
                                         AND c.cod_compositor = fc.cod_compositor
                WHERE c.nome LIKE _nome_comp
            )                                                                                                                                                               
            GROUP BY cod_grav, g.nome                                                                                                                                       
        ) AS gp                                                                                                                                                             
        GROUP BY gp.nome, gp.qtd                                                                                                                                            
    ) AS max_playlist_gravadora;
END
$$ LANGUAGE PLPGSQL;



/*
Listar nome do compositor com maior número de faixas nas playlists
existentes.
*/
CREATE OR REPLACE FUNCTION compositor_mais_nro_faixas_playlists ()
RETURNS TABLE(compositor VARCHAR)
AS $$
BEGIN
    RETURN query
    SELECT nome FROM (
        SELECT nome, MAX(qtd) FROM (
            SELECT c.nome, COUNT(distinct cod_play) qtd
            FROM compositor c
            JOIN faixa_compositor   USING(cod_compositor)
            JOIN faixa              USING(cod_faixa, cod_alb, cod_meio)
            JOIN faixa_playlist     USING(cod_faixa, cod_alb, cod_meio)
            GROUP BY c.cod_compositor, c.nome
            ) AS comp_qtd_playlists
        GROUP BY comp_qtd_playlists.nome
    ) AS compositor_mais_faixas;
END
$$ LANGUAGE PLPGSQL;

/*
SELECT nome FROM (
    SELECT nome, MAX(qtd) FROM (
        SELECT nome, COUNT(distinct cod_play) qtd                        
        FROM compositor                                                  
        JOIN faixa_compositor   USING (cod_compositor)                   
        JOIN faixa_playlist     USING (cod_faixa, cod_alb, cod_meio)     
        GROUP BY cod_compositor, nome                                    
    ) AS comp_playlists                                                  
    GROUP BY nome                                                        
) AS compositor_mais_faixas;
*/


/*
Listar playlists, cujas faixas (todas) têm tipo de composição “Concerto” e
período “Barroco”.
*/
CREATE OR REPLACE FUNCTION playlists_faixas_filtro
    (_tipo_comp composicao.descricao%TYPE DEFAULT 'Concerto', _periodo periodo_musical.periodo%TYPE DEFAULT 'barroco')
RETURNS TABLE(nome VARCHAR)
AS $$
BEGIN
    RETURN query
    SELECT p.nome FROM playlist p
    JOIN faixa_playlist     USING(cod_play)
    JOIN faixa              USING(cod_faixa, cod_alb, cod_meio)
    JOIN composicao c       USING(cod_composicao)
    JOIN faixa_compositor   USING(cod_faixa, cod_alb, cod_meio)
    JOIN compositor         USING(cod_compositor)
    JOIN periodo_musical pm USING(cod_pm)
    WHERE c.descricao = _tipo_comp AND pm.periodo = _periodo;
END
$$ LANGUAGE PLPGSQL;

