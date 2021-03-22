-- Arquivo de configuração a la bash.rc

ansikit = require 'ansikit'

prompt(ansikit.text('> {bold}'..os.getenv('USER')..': {reset} '))
