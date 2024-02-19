import requests

url = "https://localhost:8443/json"

response = requests.get(url, verify=False)

f = open("fingerprint.json", "ab")
f.write(response.content)
f.close()

print(response.text)
