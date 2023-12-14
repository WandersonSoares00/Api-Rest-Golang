
/*
Visão materializada que tem como atributos o nome da playlist e a
quantidade de álbuns que a compõem.
*/

CREATE MATERIALIZED VIEW albuns_playlist AS
SELECT nome "Playlist", COUNT(DISTINCT cod_alb) "Qtd albuns" FROM playlist
LEFT JOIN faixa_playlist USING (cod_play)
GROUP BY cod_play, nome
ORDER BY "Qtd albuns" DESC
WITH NO DATA;


