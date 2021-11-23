drop table if exists posts cascade;
drop table if exists comments;

-- テーブルの作成はgormが行う
-- ただし、テーブルの初期化はしておかなければならない
