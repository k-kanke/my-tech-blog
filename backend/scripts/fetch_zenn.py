import feedparser
import requests
import os
from datetime import datetime

API_URL = os.getenv("API_URL")
FEED_URL = os.getenv("ZENN_FEED_URL")

print(f"Fetching RSS feed: {FEED_URL}")
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

    print(f"Posting: {title} ({url})")

    try: 
        # すでにURLが存在していたらskip
        check_response = requests.get(API_URL, timeout=5)
        if check_response.status_code == 200:
            existing_articles = check_response.json()
            urls = [a['url'] for a in existing_articles]

            if url in urls:
                print("すでに存在しているのでスキップします")
                continue
        
        # 新規投稿
        response = requests.post(API_URL, json=data, timeout=10)
        print(f"POST status: {response.status_code}")
    
    except Exception as e:
        print(f"Error: {e}")
