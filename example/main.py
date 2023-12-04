
import requests
from http import HTTPStatus
import json

host = "localhost"
port = "8080"

def show_entity (entity):
    req = requests.get(f"http://{host}:{port}/api/v1/{entity}")
    res = json.loads(req.content)
    format_res = json.dumps(res, indent=2, ensure_ascii=False)
    print(f"  -=-=-=--==-=-=-=-=- Lista de {entity}  -=-=-=--==-=-=-=-=-  ")
    print(format_res)

def get_entity_by_id (entity, id):
    req = requests.get(f"http://{host}:{port}/api/v1/{entity}/{id}")
    return json.loads(req.content)

def show_entity_by_id (entity, id):
    res = get_entity_by_id(entity, id)
    format_res = json.dumps(res, indent=2, ensure_ascii=False)
    print(f"  -=-=-=--==-=-=-=-=- Lista de {entity}  -=-=-=--==-=-=-=-=-  ")
    print(format_res)

def insert_data(entity, data):
    req = requests.post(f"http://{host}:{port}/api/v1/{entity}", json=data)
    return req.status_code == HTTPStatus.CREATED

def remove_data(entity, data):
    req = requests.delete(f"http://{host}:{port}/api/v1/{entity}/{id}", json=data)
    return req.status_code == HTTPStatus.OK

def insert_faixa_in_playlist (id_play):
    while True:
        id_faixa = int(input("Informe o id da faixa a ser adicionada (-1 voltar): "))
        if id_faixa == -1:
            return True
        faixa = get_entity_by_id("faixas", id_faixa)
        
        if 'success' not in faixa and not (insert_data("faixasplaylists", {"codFaixa": id_faixa, "codAlbum": faixa['codAlbum'], "codMeio": faixa['codMeio'], "codPlay": id_play})):
            return False

def remove_faixa_from_playlist(id_play):
    while True:
        id_faixa = int(input("Informe o id da faixa a ser removida (-1 voltar): "))
        if id_faixa == -1:
            return True
        faixa = get_entity_by_id("faixas", id_faixa)
        if 'success' not in faixa and not (remove_data("faixasplaylists", {"codFaixa": id_faixa, "codAlbum": faixa['codAlbum'], "codMeio": faixa['codMeio'], "codPlay": id_play})):
            return False

def create_playlist ():
    id_play = int(input('Informe o identificador da playlist: '))
    name    = input('Informe o nome da playlist: ')

    if not insert_data("playlists", {"id": id_play, "nome": name}):
        print("Não foi possível criar a playlist")
        return
    
    while(input("Adicionar faixa? [S/n]") in ('Ss', '')):
        show_entity("albuns")
        id_album = int(input("informe o id do álbum: "))
        show_entity_by_id("faixas/albuns", id_album)
        insert_faixa_in_playlist(id_play)
        
def manage_playlists():
    id_play = int(input("Informe o id da playlist: "))
    show_entity_by_id("faixasplaylists", id_play)
    op = input("[r]emover ou [i]nserir faixas? (outro p/ voltar)")
    
    if op == 'r' and not remove_faixa_from_playlist(id_play):
        print("Não foi possível remover faixa.")
    if op == 'i' and not insert_faixa_in_playlist(id_play):
        print("Não foi possível inserir faixa.")
    else:
        return


if __name__ == '__main__':
    while True:
        show_entity("playlists")
        op = input("[c]riar ou [g]erenciar playlists ([s]air)? ")
        if op == 'c':
            create_playlist()
        if op == 'g':
            manage_playlists()
        if op == 's':
            break
