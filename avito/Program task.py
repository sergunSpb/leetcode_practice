
recom_ids = [2, 3, 1]
seen_ids = [3, 10, 20]
filtered ?= [2, 1]
# 1. надо составить новый список айдишников
# 2. в том же порядке что recom_ids
# 3. НЕ содержит seen_ids

seen_ids_dict = {x:None for x in seend_ids }
a = set(recom_ids)
filtered = []
for x in recom_ids:
    if x not in a:
        filtered.append(x)


filtered = list(set(recom_ids) - set(seen_ids))



# дан список из 1 миллиона url, которые нужно опросить
# из каждого url получить несколько item id и опросить 3 сервиса независимых c этими данными
# для из данных всех 3х сервисов собрать итоговый и ответ и сохранить результат в список
# результаты сохранить в список
from typing import List

# format: url, request_data
# sample size: 10**6
requests_samples = [
  ('http://some-service/getItems/', {'user_id': 100}),  # вернет item_ids: [1, 2, 3]
  ('http://some-service/getItems/', {'user_id': 101}),
    ...
]

service_1_url = 'http://service1/fillItems/'
service_2_url = 'http://service2/scoreItems/'
service_3_url = 'http://service3/logItems/'

def business_logic(service1_response, service2_response, service3_response):
    # эта функция не делает сетевых вызовов, только обрабатывает ответы
    # считайте, что она уже написана
    return {}

'''
# схема
1. ('http://some-service/getItems/', {'user_id': 100}) -> # вернет item_ids: [1, 2, 3]
2. в любом порядке и независимо опрашиваем service_1-2-3_url с item_ids
3. собираем результат при помощи business_logic
4. вы великолепны! повторить миллион раз.
'''


import aiohttp
import asyncio


sem = asyncio.Semaphof(1000)


def gather_data(requests: List[str]) -> List[dict]:
    res = []
    for i in requests:
       coros = [ process_url(s) for x in requests]
       res.extend(await asyncio.gather( coros ))
    return res 
    

async def process_url(url:str) ->List:
    await with sem.lock():
        with s as aiohttp.client():    
            res = await s.get(url)
            data = await res.json()
            retval = []
            cors = []
            for x in data['item_ids']:
                cors.append(helper(x))
            async for x in await asyncio.gather(cors):
                retval.append(x) 
            return retval   


async def helper(x) -> List:
    cors = []
    cors.append (s.post(f"{service_1_url}{x}") )
    cors.append (s.post(f"{service_2_url}{x}") )
    cors.append (s.post(f"{service_3_url}{x}") )
    res = await asyncio.gather( *cors )
    return business_logic(await cors[0].result.text(), await cors[1].result.text(), await cors[2].result.text())
    
                