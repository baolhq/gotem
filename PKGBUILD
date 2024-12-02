# Maintainer: baolhq <baolhq280@gmail.com>
# Contributor: Quoc Bao

pkgname=gotem
pkgver=0.0.1
pkgrel=1
pkgdesc="Go-based Tool for Efficient Management"
arch=('x86_64')
url="https://github.com/baolhq/gotem"
license=('MIT')
depends=('go' 'git')
makedepends=('go')
source=("git+https://github.com/baolhq/gotem.git" "man/gotem.1")
sha256sums=('SKIP')

build() {
  cd "$srcdir/gotem"
  
  # Set Go environment variables (if necessary, can be skipped if Go is already configured)
  export GOPATH="$srcdir/go"
  export GOROOT=/usr/lib/go
  export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

  # Build the application
  go build -o "$pkgname" .
}

package() {
  cd "$srcdir/gotem"

  # Install the binary to the correct directory
  install -Dm755 "$srcdir/gotem/$pkgname" "$pkgdir/usr/bin/$pkgname"

  # Install the man page
  install -Dm644 "$srcdir/man/gotem.1" "$pkgdir/usr/share/man/man1/gotem.1"
}

check() {
  cd "$srcdir/gotem"
  go test ./...
}

