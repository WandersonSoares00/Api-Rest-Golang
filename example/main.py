
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
    req = requests.delete(f"http://{host}:{port}/api/v1/{entity}", json=data)
    return req.status_code == HTTPStatus.OK

def insert_faixa_in_playlist (id_album, id_play):
    while True:
        id_faixa = int(input("Informe o número da faixa (-1 voltar): "))
        if id_faixa == -1:
            return True
        meio     =     input("Informe o meio físico da faixa: ")
        
        if not (insert_data(f"playlists/{id_play}/faixas", {"nroFaixa": id_faixa, "codAlbum": id_album, "meio": meio, "codPlay": id_play})):
            return False

def remove_faixa_from_playlist(id_album, id_play):
    while True:
        id_faixa = int(input("Informe o id da faixa a ser removida (-1 voltar): "))
        if id_faixa == -1:
            return True
        meio     =     input("Informe o meio físico da faixa: ")

        if not (remove_data(f"playlists/{id_play}/faixas", {"nroFaixa": id_faixa, "codAlbum": id_album, "meio": meio, "codPlay": id_play})):
            return False

def show_albuns_faixas():
    show_entity("albuns")
    id_album = int(input("informe o id do álbum: "))
    show_entity(f"albuns/{id_album}/faixas")
    return id_album

def red_str(str):
    return '\033[31m'+str+'\033[0;0m'

def create_playlist ():
    id_play = int(input('Informe o identificador da playlist: '))
    name    = input('Informe o nome da playlist: ')

    if not insert_data("playlists", {"id": id_play, "nome": name}):
        print(red_str("Não foi possível criar a playlist"))
        return
    
    if (input("Adicionar faixa? [S/n]") in ('Ss', '')):
        id_album = show_albuns_faixas()
        insert_faixa_in_playlist(id_album, id_play)

def delete_playlist ():
    id_play = int(input('Informe o identificador da playlist: '))
    if not remove_data("playlists", {"id": id_play}):
        print(red_str("Não foi possível remover a playlist, verifique se ela ainda possui faixas"))

def manage_playlists():
    id_play = int(input("Informe o id da playlist: "))
    show_entity(f"playlists/{id_play}/faixas")
    op = input("[r]emover ou [i]nserir faixas? (outro p/ voltar)")
    
    if op == 'r':
        id_album = int(input("Informe o id do álbum: "))
        if not remove_faixa_from_playlist(id_album, id_play):
            print(red_str("Não foi possível remover faixa."))
    if op == 'i':
        id_album = show_albuns_faixas()
        if not insert_faixa_in_playlist(id_album, id_play):
            print(red_str("Não foi possível inserir faixa."))
    else:
        return


if __name__ == '__main__':
    while True:
        show_entity("playlists")
        op = input("[c]riar, [a]pagar ou [g]erenciar playlists ([s]air)? ")
        if op == 'c':
            create_playlist()
        elif op == 'a':
            delete_playlist()
        elif op == 'g':
            manage_playlists()
        elif op == 's':
            print(red_str("Saindo..."))
            break
