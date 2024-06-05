# go-mqtt-example
🌟 A simple Golang MQTT example project demonstrating basic message publishing and subscribing. 🚀

## Message broker implementing the MQTT protocol

### Install On Mac :
```
brew install mosquitto
```

🚀 referenece : [goto](https://formulae.brew.sh/formula/mosquitto)

✅ Version Check :
```
mosquitto -v
```

🏁 To start Mosquitto now and restart at login:
```
brew services start mosquitto
```

⚙️ The configuration by editing:
```
 /opt/homebrew/etc/mosquitto/mosquitto.conf
```

⚠️ if you don't want/need a background service you can just run:
```
/opt/homebrew/opt/mosquitto/sbin/mosquitto -c /opt/homebrew/etc/mosquitto/mosquitto.conf
```
