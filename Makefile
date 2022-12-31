CC = go

TARGET = masscan-online

SRCS  = *.go

$(TARGET):	all

build:
	$(CC) build -o $(TARGET) $(SRCS)

all:	clean build

clean:
	-rm -f $(TARGET) 

run:
	./$(TARGET)