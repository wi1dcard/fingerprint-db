import aiohttp
import asyncio

async def main():

    async with aiohttp.ClientSession(connector=aiohttp.TCPConnector(ssl=False)) as session:
        async with session.get('https://localhost:8443/json') as response:

            f = open("fingerprint.json", "ab")
            f.write(await response.read())
            f.close()

            print(await response.text())

asyncio.run(main())
