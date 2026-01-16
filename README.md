O que é: API para sorteio de times.

Tecnologias: Go, Gin (Router), GORM (SQLite).

Como rodar: go run main.go.

Endpoints principais: POST /shuffle e GET /match/:id.

Request exemplo para /shuffle:
```curl
curl --location 'http://localhost:8080/shuffle' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Partida Teste",
    "players": [
        "Thiago '\''Canela'\'' Silva",
        "Juninho Pernambucano Jr.",
        "Betinho Furacão",
        "Zé do Brejo",
        "Ricardinho '\''Maestro'\''",
        "Bruninho '\''Pivete'\''",
        "Marcão Paredão",
        "Léo '\''Guerreiro'\'' Santos",
        "Fabinho '\''Motorzinho'\''",
        "Dudu '\''Alegria'\''",
        "Renan '\''Muralha'\''",
        "Lucas Valente",
        "Enzo Gabriel",
        "Matheus Vicenzo",
        "Gabriel '\''Gabi'\'' Martins",
        "Rafael '\''Rafa'\'' Moreno",
        "Nicolas Ferreira",
        "Victor '\''Vico'\'' Hugo",
        "André '\''Deco'\'' Lima",
        "Caio '\''Kadu'\'' Mendes",
        "Igor '\''Igui'\'' Costa",
        "Samuel '\''Sam'\'' Rocha",
        "Pedrão"
    ],
    "players_per_team": 11
}'
```

Response exemplo para /shuffle:
```json
{
    "bench": [
        "Ricardinho 'Maestro'"
    ],
    "id": "a9455328",
    "teams": [
        {
            "name": "Time 1",
            "players": [
                "Renan 'Muralha'",
                "Matheus Vicenzo",
                "Rafael 'Rafa' Moreno",
                "Caio 'Kadu' Mendes",
                "Juninho Pernambucano Jr.",
                "Léo 'Guerreiro' Santos",
                "Thiago 'Canela' Silva",
                "Betinho Furacão",
                "Pedrão",
                "Nicolas Ferreira",
                "André 'Deco' Lima"
            ]
        },
        {
            "name": "Time 2",
            "players": [
                "Lucas Valente",
                "Victor 'Vico' Hugo",
                "Igor 'Igui' Costa",
                "Bruninho 'Pivete'",
                "Gabriel 'Gabi' Martins",
                "Samuel 'Sam' Rocha",
                "Marcão Paredão",
                "Zé do Brejo",
                "Enzo Gabriel",
                "Fabinho 'Motorzinho'",
                "Dudu 'Alegria'"
            ]
        }
    ]
}
```

Request exemplo para /match/:id:
```curl
curl --location 'http://localhost:8080/match/a9455328'
```

Response exemplo para /shuffle:
```json
{
    "date": "2026-01-15T23:30:31.190564134-03:00",
    "name": "Partida Teste",
    "teams": {
        "bench": [
            "Ricardinho 'Maestro'"
        ],
        "teams": [
            {
                "name": "Time 1",
                "players": [
                    "Renan 'Muralha'",
                    "Matheus Vicenzo",
                    "Rafael 'Rafa' Moreno",
                    "Caio 'Kadu' Mendes",
                    "Juninho Pernambucano Jr.",
                    "Léo 'Guerreiro' Santos",
                    "Thiago 'Canela' Silva",
                    "Betinho Furacão",
                    "Pedrão",
                    "Nicolas Ferreira",
                    "André 'Deco' Lima"
                ]
            },
            {
                "name": "Time 2",
                "players": [
                    "Lucas Valente",
                    "Victor 'Vico' Hugo",
                    "Igor 'Igui' Costa",
                    "Bruninho 'Pivete'",
                    "Gabriel 'Gabi' Martins",
                    "Samuel 'Sam' Rocha",
                    "Marcão Paredão",
                    "Zé do Brejo",
                    "Enzo Gabriel",
                    "Fabinho 'Motorzinho'",
                    "Dudu 'Alegria'"
                ]
            }
        ]
    }
}
```