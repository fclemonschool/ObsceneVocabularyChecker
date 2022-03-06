import os

from hstest import StageTest, CheckResult, WrongAnswer, TestCase

inputs = [
    "You can’t be a friendly neighborhood spider-man if there is no neighborhood.",
    "Sometimes to do what is right we have to give up the things we want the most even our dreams.",
    "Every day I wake up knowing that the more people I try to save. The more enemies I will make.",
]

FILENAME = "quote.txt"


class TestAdmissionProcedure(StageTest):
    def generate(self):
        return [TestCase(stdin=[test], attach=[test]) for test in inputs]

    def check(self, reply: str, attach: list):
        if not os.path.exists(FILENAME):
            raise WrongAnswer("Cannot file test.txt file")

        with open(FILENAME, "r") as f:
            content = f.read().strip()
            if content != attach[0]:
                raise WrongAnswer(
                    f'Invalid content of {FILENAME} file, got "{content}" want "{attach[0]}"'
                )

        return CheckResult.correct()


if __name__ == '__main__':
    TestAdmissionProcedure().run_tests()
