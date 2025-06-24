import { useEffect, useState } from 'react';
import { fetchArticles } from '../api/articles';
import { fetchZennArticles } from '../api/zennArticles';

type CombinedArticle = {
  type: 'article' | 'zenn';
  id: number;
  title: string;
  url?: string;
  content?: string;
  date: string;
};

const Homepage = () => {
  const [combinedArticles, setCombinedArticles] = useState<CombinedArticle[]>([]);

  useEffect(() => {
    Promise.all([fetchArticles(), fetchZennArticles()]).then(([articlesRes, zennRes]) => {
      const articles: CombinedArticle[] = articlesRes.data.map((a: any) => ({
        type: 'article',
        id: a.id,
        title: a.title,
        content: a.content,
        date: a.created_at,
      }));

      const zennArticles: CombinedArticle[] = zennRes.data.map((z: any) => ({
        type: 'zenn',
        id: z.id,
        title: z.title,
        url: z.url,
        date: z.published_at,
      }));

      const combined = [...articles, ...zennArticles].sort((a, b) =>
        a.date < b.date ? 1 : -1
      );

      setCombinedArticles(combined);
    });
  }, []);

  return (
    <div style={{ padding: '20px' }}>
      <h1>ホーム - 全記事一覧</h1>
      {combinedArticles.map((item) => (
        <div key={`${item.type}-${item.id}`} style={{ marginBottom: '10px', borderBottom: '1px solid #ccc' }}>
          {item.type === 'article' ? (
            <>
              <h2>{item.title}</h2>
              <p>{item.content}</p>
              <small>{item.date}</small>
            </>
          ) : (
            <>
              <h2>
                <a href={item.url} target="_blank" rel="noopener noreferrer">
                  {item.title}
                </a>
              </h2>
              <small>{item.date}</small>
            </>
          )}
        </div>
      ))}
    </div>
  );
};

export default Homepage;
