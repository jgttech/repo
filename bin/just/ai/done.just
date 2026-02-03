# Invokes AI to generate a commit and then pushes commit.
[group('ai')]
done:
  just claude save
  git push
