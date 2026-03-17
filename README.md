this is just a websocket implementation between two programs (using [gorilla/websocket](https://github.com/gorilla/websocket). router is the server and sdlmenu is just a fancy client.


PREREQISITES:\
    - gnu cpp compiler (ive been using the msys 64 compiler)\
    - go-gl\
    - go-sdl\
    - SDL2 something something\

### SDLMENU: 
funny visual program that uses sdl. building on windows might be difficult because you need -tags static to build and that is not in the make script. the program is also a websocket client (which can run without the server and with the server) so that is why gorilla websocket is a dependency.

### ROUTER: 
funny websocket server which doesnt use sdl so its pretty much a cli program, takes in the standard input and sends it to the SDLMENU if it is available. was meant to be a complicated program that as a funny visual attachment but i stopped working on it.

### INSTRUCTIONS:
to be able to run first you have to build

to build:
`make build`  

to make specifically sdl menu or web sebsocket server:  
`make build-sdlmenu` or   
`make build-router`  
respectively   

to run both:   
`make run`    

to run one or the other:   
`make run-sdlmenu` or   
`make run-router`  


# if you find any issues make a gh issue, ill look into it asap

---

`go build -tags static`
to compile (on windows i guess)
`make build`
to compile on linux cause i dont use windows anymore lmao



