
PREFIX		=	${HOME}/.local
SHAREDIR	=	${PREFIX}/share/edpad2
CONFDIR		=	${PREFIX}/etc
GOBIN		=	${PREFIX}/bin

GO111MODULE	=	auto

all: build

build:
	go build ./cmd/edpad2

install: build
	go env -w GOBIN=${GOBIN}
	go install ./cmd/edpad2
	mkdir -p ${SHAREDIR}
	cp -f res/edpad2.css ${SHAREDIR}/edpad2.css
	sed -i "s=@SHAREDIR@=${SHAREDIR}=g" ${SHAREDIR}/edpad2.css
	cp -f res/edpad2.ui ${SHAREDIR}
	cp -f res/*.wav ${SHAREDIR}
	cp -f res/edpad2.png ${SHAREDIR}/edpad2.png
	mkdir -p ${CONFDIR}
	cp -r etc/edpad2.conf ${CONFDIR}/edpad2.conf

