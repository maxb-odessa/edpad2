
PREFIX		=	${HOME}/.local
SHAREDIR	=	${PREFIX}/share/edpad2
FONTDIR		=	${PREFIX}/share/fonts
CONFDIR		=	${PREFIX}/etc
GOBIN		=	${PREFIX}/bin

GO111MODULE	=	auto

all: build

build:
	go build ./cmd/edpad2

install:
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
	mkdir -p ${FONTDIR}
	cp -f res/*.ttf ${FONTDIR}
	fc-cache -f

