# quartz
golang file watch

## what can it do

trace adding file and updating file

## how to use

```golang
q, err := NewQuartz("./test/", time.Second)
	if err != nil {
		panic(err)
	}
q.Begin()
defer q.Stop()
select {
case event:=<-q.Event:
	log.Println(event)
}
```

## what to do
- platform fs notify
- zan by mq

