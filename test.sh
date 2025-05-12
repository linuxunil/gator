echo "reset"
go run . reset
echo "register"
go run . register lane
echo "add"
go run . addfeed "TechCrunch:" "https://techcrunch.com/feed/"
go run . addfeed "Boot.dev Blog:" "https://blog.boot.dev/index.xml"
go run . addfeed "Hacker News:" "https://news.ycombinator.com/rss"
echo "feeds"
go run . feeds
