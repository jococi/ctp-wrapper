import platform

version = "0.0.1"

if "Windows" in platform.system():
    from pyctp.windows import *
elif "Linux" in platform.system():
    from pyctp.linux import *
elif "Darwin" in platform.system():
    from pyctp.macos import *
