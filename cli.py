
import fire
import shutil


def to_camel_case(snake_str) -> str:
    components = snake_str.split('_')
    return components[0] + ''.join(x.title() for x in components[1:])


qut_object_name = "'object'"
object_name = 'object'
camel = 'object'
up_camel = 'Object'


class Hello:
    def __init__(self):
        self.x = 1
        self.qut_object_name = None
        self.object_name = None
        self.camel = None
        self.up_camel = None

    def create_template(self, object_name, placeholder):
        self.object_name = object_name
        self.qut_object_name = "'" + self.object_name + "'"
        self.camel = to_camel_case(object_name)
        self.up_camel = self.camel.title()
        print("hello cli", self.object_name, self.camel, self.up_camel, placeholder)

        object_file = object_name + '.go'

        self.template_to('./model/db-layer/object.go', object_file)

    def template_to(self, src, dst):
        shutil.copyfile(src, dst)
        with open(dst, 'r+') as obj_template:
            source = obj_template.read()
            source = source.replace(qut_object_name, self.qut_object_name)
            source = source.replace(camel, self.camel)
            source = source.replace(up_camel, self.up_camel)
            # print(source)

            obj_template.seek(0)
            obj_template.write(source)


if __name__ == '__main__':
    fire.Fire(Hello)

