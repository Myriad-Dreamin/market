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
        if not cls.silent: print(cmd)
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
    @require_cls('read_context')
    def image(cls, *_):
        cls.pipe('docker build --tag %s .' % cls.context['node-name'])
    
    @classmethod
    def run_testdb(cls, *_):
        cls.pipe('docker run -id -p 23306:3306 -e MYSQL_ROOT_PASSWORD=12345678 -e MYSQL_DATABASE=market -e MYSQL_USER=madmin -e MYSQL_PASSWORD=12345678 --name backend-testdb mysql:5.7')

    @classmethod
    def stop_testdb(cls, *_):
        cls.pipe('docker stop backend-testdb')

    @classmethod
    def start_testdb(cls, *_):
        cls.pipe('docker start backend-testdb')

    @classmethod
    def remove_testdb(cls, *_):
        cls.pipe('docker rm backend-testdb')

    @classmethod
    @require_cls('read_context')
    def run(cls, config_file=None, *_):
        if config_file is not None:
            cls.config_file = os.path.abspath(config_file)
            cls.config_file_target = os.path.join("/", os.path.basename(config_file))
        cls.pipe('docker run -id -v %s:%s --network=host --name %s %s' %\
            (cls.config_file, cls.config_file_target,
            cls.context['instance-name'], cls.context['node-name']))
    
    @classmethod
    @require_cls('read_context')
    def start_run(cls, *_):
        cls.pipe('docker start %s' % cls.context['instance-name'])

    @classmethod
    @require_cls('read_context')
    def stop_run(cls, *_):
        cls.pipe('docker stop %s' % cls.context['instance-name'])

    @classmethod
    @require_cls('read_context')
    def remove_run(cls, *_):
        cls.pipe('docker rm %s' % cls.context['instance-name'])

    @classmethod
    @require_cls('template')
    def up(cls, *_):
	    pipe('docker-compose -f %s up' % (cls.compose_file)) 

    @classmethod
    @require_cls('template')
    def down(cls, *_):
	    pipe('docker-compose -f %s down' % (cls.compose_file)) 
        
    @classmethod
    @require_cls('template')
    def start(cls, *_):
	    pipe('docker-compose -f %s start' % (cls.compose_file)) 
        
    @classmethod
    @require_cls('template')
    def stop(cls, *_):
	    pipe('docker-compose -f %s stop' % (cls.compose_file)) 

    @classmethod
    @require_cls('read_context')
    def template(cls, *_):
        with open(cls.compose_template_file) as f:
            s = f.read().replace('{{redis-root-password}}', cls.context['redis-root-password'])
            s = s.replace('{{mysql-root-password}}', cls.context['mysql-root-password'])
            s = s.replace('{{conf-path}}', cls.context['conf-path'])
            s = s.replace('{{logs-path}}', cls.context['logs-path'])
            s = s.replace('{{data-path}}', cls.context['data-path'])
            s = s.replace('{{node-name}}', cls.context['node-name'])
            s = s.replace('{{instance-name}}', cls.context['instance-name'])
            s = s.replace('{{target-port}}', str(cls.context['target-port']))
            s = s.replace('{{config-file}}', cls.config_file)
            s = s.replace('{{config-file-target}}', cls.config_file_target)
            s = s.replace('{{mysql-norm-password}}', cls.context['mysql-norm-password'])
            with open(cls.compose_file, 'w') as o:
                o.write(s)

    @classmethod
    def read_context(cls, *_):
        with open('.market-env.json', 'w+') as f:
            c = f.read().encode('utf-8')
            cls.context = json.loads('{}' if len(c) == 0 else c)
            cls.context['node-name'] = cls.context.get('node-name', 'minimum-market/backend:alpine')
            cls.context['instance-name'] = cls.context.get('instance-name', 'backend')
            
            cls.context['redis-root-password'] = cls.context.get('redis-root-password', '12345678')
            cls.context['mysql-root-password'] = cls.context.get('mysql-root-password', '12345678')
            cls.context['mysql-norm-password'] = cls.context.get('mysql-norm-password', '12345678')
            cls.context['conf-path'] = os.path.join(cls.current_path, cls.context.get('conf-path', 'testdb/conf'))
            cls.context['logs-path'] = os.path.join(cls.current_path, cls.context.get('logs-path', 'testdb/logs'))
            cls.context['data-path'] = os.path.join(cls.current_path, cls.context.get('data-path', 'testdb/data'))

            cls.context['target-port'] = cls.context.get('target-port', 23335)
            
            cls.config_file_name = cls.context.get('config-file-name', 'config.toml')
            cls.config_file = os.path.join(cls.current_path, cls.config_file_name)
            cls.config_file_target = os.path.join('/', cls.config_file_name)

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
