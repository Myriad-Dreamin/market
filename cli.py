
import fire


class Hello:
    def __init__(self):
        self.x = 1

    def create_template(self, object_name, placeholder):
        print("hello cli", object_name, placeholder)


if __name__ == '__main__':
    fire.Fire(Hello)

