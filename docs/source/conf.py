# Configuration file for the Sphinx documentation builder.
# Read https://www.sphinx-doc.org/en/master/usage/configuration.html for more options available

# import sphinx_pdj_theme

# -- Project information

project = 'IP2Location Go'
copyright = '2023, IP2Location'
author = 'IP2Location'

release = '1.0.0'
version = '1.0.0'

# -- General configuration

extensions = [
    'sphinx.ext.duration',
    'sphinx.ext.doctest',
    'myst_parser',
    'sphinx_copybutton',
]

# https://myst-parser.readthedocs.io/en/latest/configuration.html#setting-html-metadata
myst_html_meta = {
    "description": "IP2Location Go module provides a fast lookup of geolocation information from IP address by using IP2Location database.",
    "keywords": "IP2Location, Geolocation, IP location, Go",
    "google-site-verification": "DeW6mXDyMnMt4i61ZJBNuoADPimo5266DKob7Z7d6i4",
}

# https://myst-parser.readthedocs.io/en/latest/syntax/optional.html

myst_enable_extensions = [
    "colon_fence",
    "deflist",
    "fieldlist",
]

# templates_path = ['_templates']

# -- Options for HTML output

html_theme = 'sphinx_book_theme'
# html_theme_path = [sphinx_pdj_theme.get_html_theme_path()]

# PDJ theme options, see the list of available options here: https://github.com/jucacrispim/sphinx_pdj_theme/blob/master/sphinx_pdj_theme/theme.conf
html_theme_options = {
    "use_edit_page_button": False,
    "use_source_button": False,
    "use_issues_button": False,
    "use_download_button": False,
    "use_sidenotes": False,
}

# The name of an image file (relative to this directory) to place at the top
# of the sidebar.
html_logo = 'images/ipl-logo-square-1200.png'

# Favicon
html_favicon = 'images/favicon.ico'

html_title = "IP2Location Go"

# html_baseurl = "https://ip2location-go.readthedocs.io/en/latest/"