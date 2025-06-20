import feedparser
import requests
import os
from datetime import datetime

API_URL = os.getenv("API_URL")
FEED_URL = os.getenv("ZENN_FEED_URL")

feed = feedparser.parse(FEED_URL)

for entry in feed.entries:
    title = entry.title
    url = entry.link
    published_at = entry.published_parsed
    published_at_iso = datetime(*published_at[:6]).isoformat() + "Z"

    data = {
        "title": title,
        "url": url,
        "published_at": published_at_iso
    }

    response = requests.post(API_URL, json=data)
    print(f"POST {url} → {response.status_code}")
