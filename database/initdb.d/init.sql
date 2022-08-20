DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(100), 
  time REAL NOT NULL, 
  image_url TEXT NOT NULL, 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO posts(name, time, image_url, created_at, updated_at) VALUES
  ("alice", 2.0, "https://ja.wikipedia.org/wiki/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB:%E7%B7%9A%E9%A6%99%E8%8A%B1%E7%81%AB.jpg#/media/ファイル:線香花火.jpg", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ("bob", 3.0, "https://news.1242.com/wp-content/uploads/2021/07/5074083_sRS.jpg", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
;
