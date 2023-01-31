export class GitRepo {
    name: string = ""
    repoURL: string = ""

    test() {
        return "Hello, " + this.name + " - " + this.repoURL;
    }
}
