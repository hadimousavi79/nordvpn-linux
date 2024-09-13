import pytest
import sh

import lib
from lib import (
    daemon,
    info,
    logging,
    login,
    network,
)

def setup_function(function):  # noqa: ARG001
    logging.log()


def teardown_function(function):  # noqa: ARG001
    logging.log(data=info.collect())
    logging.log()


def test_loop_50():
    for x in range(50):
        daemon.start()
        daemon.stop()
