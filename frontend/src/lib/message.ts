export class Message {
    nickname: string;
    text: string;
    created_at: Date;

    constructor(nickname: string, text: string, created_at: Date) {
        this.nickname = nickname;
        this.text = text;
        this.created_at = created_at;
    }
}
