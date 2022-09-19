# 開発環境構築方法

1. Dockerイメージの作成

   ```shell
   sudo docker compose up -d --build
   ```

2. 作られたDockerイメージの確認

   ```shell
   sudo docker compose ps
   ```

3. コンテナへの接続

   ```shell
   sudo docker compose exec pythondev bash
   ```

   

