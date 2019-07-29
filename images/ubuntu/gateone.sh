#!/bin/sh

apt -y install libkrb5-dev
pip install tornado kerberos;
git clone https://github.com/liftoff/GateOne.git;
cd GateOne;
python setup.py install

pip uninstall -y tornado
pip install tornado==4.5.3
