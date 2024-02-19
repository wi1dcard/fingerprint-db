import httpx

url = "https://localhost:8443/json"

client = httpx.Client(http2=True, verify=False)
response = client.get(url)

f = open("fingerprint.json", "ab")
f.write(response.content)
f.close()

print(response.text)
