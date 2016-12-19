# Resume for Homin Lee

* Live at [here](http://suapapa.github.io/resume/).
* Based on [philipithomas/jekyll-resume](https://github.com/philipithomas/jekyll-resume/).

# Testrun

Run jekyll/jekyll image:

    $ docker run -v $(pwd):/srv/jekyll -p 80:4000 -it jekyll/jekyll /bin/bash

Test run:

    # jekyll s -b ""

Build:

    # jekyll b

# Deploy

    $ cd _site
    $ git init . && git add . && git commit -m "built at 2016-12-19"
    $ git push -f https://github.com/suapapa/resume master:gh-pages
