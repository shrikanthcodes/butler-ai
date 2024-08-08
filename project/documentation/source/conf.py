# Configuration file for the Sphinx documentation builder.
#
# For the full list of built-in configuration values, see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Project information -----------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#project-information

import os
import sys
sys.path.insert(0, os.path.abspath('../../src'))

try:
    import utils.formatting
    print("Module imported successfully")
except ImportError as e:
    print(f"Error importing module: {e}")


project = 'Butler.ai Documentation'
copyright = '2024, Shrikanth Subramanian'
author = 'Shrikanth Subramanian, Hithesh Duttuluri'
release = '0.2'

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = [
    'sphinx.ext.autodoc',
    'sphinx.ext.napoleon',
    'sphinx.ext.viewcode'
]

# -- Options for HTML output -------------------------------------------------
html_theme = 'sphinx_rtd_theme'

# -- Options for HTML output -------------------------------------------------
templates_path = ['_templates']
exclude_patterns = []

# -- Options for HTML output -------------------------------------------------
language = 'English'

# -- Options for HTML output -------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#options-for-html-output

html_theme = 'alabaster'
html_static_path = ['_static']
