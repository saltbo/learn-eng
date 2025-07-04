import Crawler from "crawler";

import fs from "fs";

const cookieStr = fs.readFileSync("cookie.txt", "utf8");
console.log(cookieStr);

const c = new Crawler({
  maxConnections: 10,
  rateLimit: 1000,
  headers: {
    "User-Agent":
      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.101.76 Safari/537.36",
    Fromurl: "ieltscat.xdf.cn",
    Cookie: cookieStr,
  },
  jQuery: false,
  // This will be called for each crawled page
  callback: (error, res, done) => {
    if (error) {
      console.log(error);
    } else {
      console.log();
      //   const $ = res.$;
      // $ is Cheerio by default
      //a lean implementation of core jQuery designed specifically for the server
      //   console.log($("title").text());
    }
    done();
  },
});

const timestamp = Date.now();
console.log(timestamp);
c.add([
  {
    url: `https://ieltscat.xdf.cn/api/list/login/ielts/2/order?_t=${timestamp}&value=0`,
    jQuery: false,
    callback: (error, res, done) => {
      if (error) {
        console.log(error);
      } else {
        console.log(res.body);
        const data = JSON.parse(res.body).data;
        console.log(data.jianyaList);

        for (let i = 0; i < data.jianyaList.length; i++) {
          console.log(data.jianyaList[i].jianyaId);
          c.add([
            {
              url: `https://ieltscat.xdf.cn/api/list/login/ielts/2/order?_t=${timestamp}&value=${data.jianyaList[i].jianyaId}`,
              jQuery: false,
              callback: (error, res, done) => {
                if (error) {
                  console.log(error);
                } else {
                  console.log(res.body);
                  const data = JSON.parse(res.body).data;
                  for (let i = 0; i < data.length; i++) {
                    const name = data[i].name;
                    const sectionList = data[i].sectionList;
                    // console.log(name, sectionList);

                    for (let j = 0; j < sectionList.length; j++) {
                      console.log(sectionList[j].questionId);
                      c.add([
                        {
                          url: `https://ieltscat.xdf.cn/api/questionPreviewQuestion/${sectionList[j].questionId}?_t=${timestamp}`,
                          jQuery: false,
                          callback: (error, res, done) => {
                            if (error) {
                              console.log(error);
                            } else {
                              console.log(res.body);
                              const { number, article, qTopic, docList } =
                                JSON.parse(res.body).data;
                              const content = docList[1].list[0].body;
                              console.log(content);
                              fs.writeFileSync(
                                `./${number}.txt`,
                                JSON.stringify({
                                  number,
                                  topic: qTopic,
                                  title: article,
                                  content,
                                })
                              );
                            }
                            done();
                          },
                        },
                      ]);
                    }
                  }

                  //   for (let i = 0; i < sectionList.length; i++) {
                  //     console.log(sectionList[i].questionId);
                  //   }
                }
                done();
              },
            },
          ]);
        }

        done();
      }
    },
  },
]);
c.add(
  `https://ieltscat.xdf.cn/api/list/login/ielts/2/order?_t=${timestamp}&value=0`
);
