targets=yj jy
sources=main.go

all: $(targets)

jy: yj
	ln -s yj jy

yj: $(sources)
	go build -o $@ $^

clean:
	rm -f yj jy
