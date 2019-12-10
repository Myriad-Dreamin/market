#!/usr/bin/env python3
import os, re, sys, shutil, json
from pymake import require_cls, oqs, entry, pipe

class Makefile:
    current_path = os.path.dirname(os.path.realpath(__file__))
    config_file_name = 'config.toml'
    context = dict()
    
    compose_template_file = 'docker-compose.template.yml'
    compose_file = 'docker-compose.yml'
    silent = False
    
    @classmethod
    def hello(cls, *_):
        print('minimum builder')

    @classmethod
    def pipe(cls, cmd, *_):
        if not Makefile.silent: print(cmd)
        pipe(cmd)
    
    @classmethod
    def generate(cls, path='./', match=None, *_):
        match = cls._gen_match(match)
        for file in os.listdir(path):
            file = os.path.join(path, file)
            if os.path.isdir(file):
                cls.generate(file, match)
            if os.path.isfile(file) and match.match(file):
                cls.pipe('go generate %s' % file)

    @staticmethod
    def _gen_match(match):
        if match is None:
            match = re.compile(r'^.*\.go$')
        if isinstance(match, str):
            match = re.compile(match)
        return match

    @classmethod
    def image(cls, *_):
        cls.pipe('docker build --tag %s .' % Makefile.context['node-name'])
        
    @classmethod
    @require_cls('template')
    def up(cls, *_):
	    pipe('docker-compose -f %s up' % (Makefile.compose_file)) 

    @classmethod
    @require_cls('template')
    def down(cls, *_):
	    pipe('docker-compose -f %s down' % (Makefile.compose_file)) 
        
    @classmethod
    @require_cls('template')
    def start(cls, *_):
	    pipe('docker-compose -f %s start' % (Makefile.compose_file)) 
        
    @classmethod
    @require_cls('template')
    def stop(cls, *_):
	    pipe('docker-compose -f %s stop' % (Makefile.compose_file)) 

    @classmethod
    @require_cls('read_context')
    def template(cls, *_):
        with open(Makefile.compose_template_file) as f:
            s = f.read().replace('{{redis-root-password}}', Makefile.context['redis-root-password'])
            s = s.replace('{{mysql-root-password}}', Makefile.context['mysql-root-password'])
            s = s.replace('{{conf-path}}', Makefile.context['conf-path'])
            s = s.replace('{{logs-path}}', Makefile.context['logs-path'])
            s = s.replace('{{data-path}}', Makefile.context['data-path'])
            s = s.replace('{{node-name}}', Makefile.context['node-name'])
            s = s.replace('{{instance-name}}', Makefile.context['instance-name'])
            s = s.replace('{{target-port}}', str(Makefile.context['target-port']))
            s = s.replace('{{config-file}}', Makefile.config_file)
            s = s.replace('{{config-file-target}}', Makefile.config_file_target)
            s = s.replace('{{mysql-norm-password}}', Makefile.context['mysql-norm-password'])
            with open(Makefile.compose_file, 'w') as o:
                o.write(s)

    @classmethod
    def read_context(cls, *_):
        with open('.market-env.json', 'w+') as f:
            c = f.read().encode('utf-8')
            Makefile.context = json.loads('{}' if len(c) == 0 else c)
            Makefile.context['node-name'] = Makefile.context.get('node-name', 'minimum-market/backend:alpine')
            Makefile.context['instance-name'] = Makefile.context.get('instance-name', 'backend')
            
            Makefile.context['redis-root-password'] = Makefile.context.get('redis-root-password', '12345678')
            Makefile.context['mysql-root-password'] = Makefile.context.get('mysql-root-password', '12345678')
            Makefile.context['mysql-norm-password'] = Makefile.context.get('mysql-norm-password', '12345678')
            Makefile.context['conf-path'] = os.path.join(Makefile.current_path, Makefile.context.get('conf-path', 'testdb/conf'))
            Makefile.context['logs-path'] = os.path.join(Makefile.current_path, Makefile.context.get('logs-path', 'testdb/logs'))
            Makefile.context['data-path'] = os.path.join(Makefile.current_path, Makefile.context.get('data-path', 'testdb/data'))

            Makefile.context['target-port'] = Makefile.context.get('target-port', 23335)
            
            Makefile.config_file_name = Makefile.context.get('config-file-name', 'config.toml')
            Makefile.config_file = os.path.join(Makefile.current_path, Makefile.config_file_name)
            Makefile.config_file_target = os.path.join('/', Makefile.config_file_name)

    @classmethod
    @require_cls('read_context')
    def clean(cls, *_):
        shutil.rmtree('testdb')

    @classmethod
    @require_cls('up')
    def all(cls, *_):
        pass

if __name__ == '__main__':
    entry(Makefile, sys.argv[1:])
