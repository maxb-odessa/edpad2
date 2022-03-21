
PREFIX		=	${HOME}/.local
SHAREDIR	=	${PREFIX}/share/edpad2
CONFDIR		=	${PREFIX}/etc
GOBIN		=	${PREFIX}/bin

GO111MODULE	=	auto

CC="gcc -march=znver3"

all: build

build:
	go build ./cmd/edpad2

install:
	go env -w GOBIN=${GOBIN}
	go install ./cmd/edpad2
	mkdir -p ${SHAREDIR}
	cp -f res/edpad2.css.in ${SHAREDIR}/edpad2.css
	sed -i "s=@SHAREDIR@=${SHAREDIR}=g" ${SHAREDIR}/edpad2.css
	cp -f res/edpad2.glade ${SHAREDIR}
	cp -f res/edpad-1280-800.png ${SHAREDIR}/edpad2.png

