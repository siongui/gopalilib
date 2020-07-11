=================================
Pāli_ Library and Data Processing
=================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/gopalilib?status.svg
   :target: https://godoc.org/github.com/siongui/gopalilib

.. image:: https://api.travis-ci.org/siongui/gopalilib.svg?branch=master
   :target: https://travis-ci.org/siongui/gopalilib

.. image:: https://goreportcard.com/badge/github.com/siongui/gopalilib
   :target: https://goreportcard.com/report/github.com/siongui/gopalilib

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/gopalilib/blob/master/UNLICENSE

This repository consist of two part:

1. Common Go_ code (library of `Pāli project`_) to be shared and used by
   client-side (front-end browser) and
   server-side (back-end server).
   The code is located in `lib <lib>`_ directory.

2. Offline data processing.

   - test_bookparser
   - test_wordparser
   - test_vfsbuild needs test_wordparser
   - test_symlink needs test_vfsbuild

UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `reflection - What are the use(s) for tags in Go? - Stack Overflow <https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go>`_
.. [2] `GitHub Pages Symbolic Link Caveat <https://siongui.github.io/2017/03/30/github-pages-symlink-caveat/>`_

.. _Go: https://golang.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _Pāli project: https://github.com/siongui/pali
.. _UNLICENSE: http://unlicense.org/
