rsync -vltzh -e ssh -r --include 'books' --exclude '*' ./ ss:/data/lxw/book/
rsync -vltzh -e ssh --delete -r $GOPATH/src/books/views/* ss:/data/lxw/book/views/
