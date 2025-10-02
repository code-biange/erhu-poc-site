#!/bin/bash

# microCMSの認証情報を環境変数から取得
SERVICE_ID=${MICROCMS_SERVICE_ID}
API_KEY=${MICROCMS_API_KEY}

# コンテンツを保存するディレクトリを作成
mkdir -p data/cms

# ヘッダー情報を設定
HEADER_KEY="X-MICROCMS-API-KEY: ${API_KEY}"

# --- APIデータを取得し、JSONファイルとして保存 ---

# 1. top_page (単一コンテンツ)
curl -s -X GET "https://${SERVICE_ID}.microcms.io/api/v1/top_page" -H "${HEADER_KEY}" > data/cms/top_page.json

# 2. news (リスト形式)
curl -s -X GET "https://${SERVICE_ID}.microcms.io/api/v1/news?limit=100" -H "${HEADER_KEY}" > data/cms/news.json

# 3. liveconcerts (リスト形式)
curl -s -X GET "https://${SERVICE_ID}.microcms.io/api/v1/liveconcerts?limit=100" -H "${HEADER_KEY}" > data/cms/liveconcerts.json

echo "microCMSデータの取得と保存が完了しました。"
