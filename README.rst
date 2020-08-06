=================================
Pāli_ Library and Data Processing
=================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/gopalilib?status.svg
   :target: https://godoc.org/github.com/siongui/gopalilib

.. image:: https://api.travis-ci.org/siongui/gopalilib.svg?branch=master
   :target: https://travis-ci.org/siongui/gopalilib

.. image:: https://gitlab.com/siongui/gopalilib/badges/master/pipeline.svg
    :target: https://gitlab.com/siongui/gopalilib/-/commits/master

.. image:: https://goreportcard.com/badge/github.com/siongui/gopalilib
   :target: https://goreportcard.com/report/github.com/siongui/gopalilib

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/gopalilib/blob/master/UNLICENSE

This repository consist of:

1. Common Go_ code (library of `Pāli project`_) to be shared and used by
   client-side (front-end browser) and
   server-side (back-end server).
   The code is located in `lib <lib>`_ directory.

2. Bootstrap/setup Dictionary (Offline data processing).
   The code is located in `dicutil <dicutil>`_.

   - test_bookparser
   - test_wordparser
   - test_triebuild needs test_wordparser
   - test_vfsbuild needs test_wordparser
   - test_symlink needs test_vfsbuild
   - test_embedmetadata needs test_bookparser and test_triebuild

2. Bootstrap/setup Tipiṭaka (Offline data processing).
   The code is located in `tpkutil <tpkutil>`_.

   - test_download_tpk
   - test_build_tpk_tree


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
