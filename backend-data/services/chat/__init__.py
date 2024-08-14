import os
import sys

# Manually set PYTHONPATH for testing purposes
sys.path.insert(0, os.path.abspath(os.path.dirname(__file__)))

# Check if PYTHONPATH is correctly set
print(f"PYTHONPATH: {os.path.abspath(os.path.dirname(__file__))}")
