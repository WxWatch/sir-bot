import { Client, TextChannel } from 'discord.js';
import { ICron } from './../interfaces';

const cron: ICron = {
    name: 'Intel',
    interval: 86400000,
    action: (client: Client) => {
        const guild = client.guilds.get('160135882274373633');
        if (!guild) {
            return;
        }
        const channel = guild.channels.get('160135882274373633') as TextChannel;
        if (!channel) {
            return;
        }
        channel.send(`PSA: Contrary to some people's beliefs, AMD Ryzen processers are not as good as Intel Core series processors`);
    }
};

export default cron;
