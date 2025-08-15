# 🌤️ Weather API

A simple API to fetch the current temperature of a location using the [OpenWeatherMap API](https://openweathermap.org/api).

---

## 🚀 Endpoint

### GET `/retrieve-temperature`

Fetches the current temperature for a specific location.

**Query Parameters:**

| Parameter | Description                 | Required |
| --------- | --------------------------- | -------- |
| city      | Name of the city            | Yes      |
| state     | State or region of the city | No       |
| country   | Country code (e.g., BR)     | No       |

**💻 Request Examples:**

```bash
curl "https://weather-api-production-944e.up.railway.app/retrieve-temperature?city=sao paulo"
```

```bash
curl "https://weather-api-production-944e.up.railway.app/retrieve-temperature?city=sao paulo&state=SP"
```

```bash
curl "https://weather-api-production-944e.up.railway.app/retrieve-temperature?city=sao paulo&state=SP&country=BR"
```

**🌡️ Response Example:**

```json
{
  "location": {
    "lat": -23.5506507,
    "lon": -46.6333824,
    "name": "São Paulo",
    "state": "São Paulo",
    "country": "BR"
  },
  "temperature": {
    "kelvin": 287.86,
    "celsius": 14.71,
    "fahrenheit": 58.48
  }
}
```

---

## ⚠️ Error Handling

If something goes wrong with the external API or the request, the server responds with the proper HTTP status code but error message only when its a Internal Server Error.

| Status Code | Meaning                       |       |  
| ----------- | ----------------------------- | ----- |
| 500         | Internal Server Error         | ❌    |
| 404         | Location Not Found (optional) | ❌    |

**💻 Example Error Response (500):**

```http
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8

Error: OpenWeatherMap API unreachable
```

---

## 📄 Notes

* `state` and `country` are optional;
* `curl` is the easiest way to test your API from the terminal. 🖥️
* Use the `run.sh` script to run the API quickly. 🚀 (api key is required)
* All temperature values are rounded to **2 decimal places**. 🔢
