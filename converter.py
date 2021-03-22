import requests

cache = {}
x = input('Enter your currency code: ').upper()
r = requests.get(f'http://www.floatrates.com/daily/{x}.json')
if x != 'USD':
    cache['USD'] = r.json()['usd']['rate']
if x != 'EUR':
    cache['EUR'] = r.json()['eur']['rate']

while True:
    y = input('Enter exchange currency: ').upper()
    if not y:
        break
    num = input(f'Enter amount of {x}: ')
    if '.' in num:
        num = float(num)
    else:
        num = int(num)
    print('Checking the cache...')
    if y in cache:
        print('Oh! It is in the cache!')
    else:
        print('Sorry, but it is not in the cache!')
        cache[y] = r.json()[y.lower()]['rate']
    print(f'You received {round(cache[y]*num, 2)} {y}')
