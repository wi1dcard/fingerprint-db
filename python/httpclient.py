import http.client
import ssl

conn = http.client.HTTPSConnection(
    "localhost:8443",
    context = ssl._create_unverified_context()
)
conn.request("GET", "/json")
res = conn.getresponse()
data = res.read()

f = open("fingerprint.json", "ab")
f.write(data)
f.close()

print(data.decode("utf-8"))
