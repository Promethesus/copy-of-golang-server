# based on a paper for dota 2 
# https://www.cs.ru.nl/bachelors-theses/2016/Wietse_Kuipers___4317904___Improving_matchmaking_with_game_data.pdf
# this script will generate player data for testing match making algorithms 

import random
import json

def generate_player_data(num_players):
    players = []
    for _ in range(num_players):
        player = {
            'last_hits': random.randint(0, 100),
            'item_5': random.randint(0, 200),
            'hero_damage': random.randint(1000, 20000),
            'gold_spent': random.randint(1000, 20000),
            'gold_per_min': random.randint(200, 500),
            'denies': random.randint(0, 10),
            'level': random.randint(1, 25),
            'item_0': random.randint(0, 200),
            'item_3': random.randint(0, 50),
            'gold': random.randint(1000, 5000),
            'deaths': random.randint(0, 10),
            'assists': random.randint(0, 30),
            'xp_per_min': random.randint(300, 600),
            'leaver_status': random.randint(0, 1),
            'account_id': random.randint(10000000, 99999999),
            'ability_upgrades': [
                {
                    'time': random.randint(100, 500),
                    'level': random.randint(1, 3),
                    'ability': random.randint(5000, 6000)
                }
            ],
            'player_slot': random.randint(0, 10),
            'hero_healing': random.randint(0, 200),
            'item_4': random.randint(0, 150),
            'item_1': random.randint(0, 10),
            'hero_id': random.randint(1, 100),
            'tower_damage': random.randint(0, 100),
            'kills': random.randint(0, 15),
            'item_2': random.randint(0, 100)
        }
        players.append(player)
    return players



with open('fake_dota_data.json', 'w') as w:
    json.dump(generate_player_data(10),w)
