# yayagf

yet another yet another go web framework

## install

`go install ./cmd/yayagf`

## what to use

Everything I write is just stealing from rails.

### performance

web framework should not treat the performance very important.

go is fast, so performance is not a really big thing.

## how to use

### yayagf new

type `yayagf new gitlab.papegames.com/fengche/abc` will generate a folder named abc

a project structure is like this

project_root
-app
--controller
--model
--ent
--serializer // TODO
--docs // TODO
-config_file1 // TODO
-config_file2 // TODO
-project_root.yml // TODO
-Dockerfile // TODO

### yayagf server

go into any go project 

`yayagf server` will monitor the code, recompile the code when compilable and run.

### yayagf generate

generate a http server scaffold.

#### router

#### serializer

#### model (ent)

#### docs (swagger) // TODO
