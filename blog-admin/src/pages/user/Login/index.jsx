import {
  LockOutlined,
  MailOutlined
} from '@ant-design/icons';
import { Alert, message } from 'antd';
import React, { useState } from 'react';
import { ProFormCheckbox, ProFormText, LoginForm } from '@ant-design/pro-form';
import { history, useModel } from 'umi';
import Footer from '@/components/Footer';
import { login } from '@/services/user';
import styles from './index.less';

const LoginMessage = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Login = () => {
  const { initialState, setInitialState } = useModel('@@initialState');
  const [userLoginState, setUserLoginState] = useState({});

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();

    if (userInfo) {
      await setInitialState((s) => ({ ...s, currentUser: userInfo }));
    }
  };

  const handleSubmit = async (values) => {
    try {
      const res = await login(values);
      if (res.success) {
        message.success('登录成功！');
        await fetchUserInfo();
        if (!history) return;
        const { query } = history.location;
        const { redirect } = query;
        history.push(redirect || '/');
        return;
      }
      console.log(msg);
      setUserLoginState(msg);
    } catch (error) {
      const { errorCode, errorMessage } = error
      message.error(`登录失败，${errorCode} ${errorMessage}`);
    }
  };


  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <LoginForm
          logo={<img alt="logo" src="/logo.svg" />}
          title="BuChiYu"
          subTitle="后台管理系统"
          initialValues={{
            autoLogin: true,
          }}
          actions={[]}
          onFinish={async (values) => await handleSubmit(values)}
        >
          {userLoginState?.errorMessage && <LoginMessage content={userLoginState?.errorMessage} />}
          <ProFormText
            fieldProps={{
              size: 'large',
              prefix: <MailOutlined className={styles.prefixIcon} />,
            }}
            name="email"
            placeholder="邮箱"
            rules={[
              {
                required: true,
                message: '请输入邮箱！',
              },
              // {
              //   pattern: /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/,
              //   message: '邮箱格式错误！',
              // },
            ]}
          />
          <ProFormText.Password
            name="password"
            fieldProps={{
              size: 'large',
              prefix: <LockOutlined className={styles.prefixIcon} />,
            }}
            placeholder="密码"
            rules={[
              {
                required: true,
                message: '请输入密码！',
              },
              // {
              //   pattern: /^(?!([a-zA-Z]+|\d+)$)[a-zA-Z\d]{6,20}$/,
              //   message: '密码长度为8-32',
              // },
            ]}
          />
          <div
            style={{
              marginBottom: 24,
            }}
          >
            <ProFormCheckbox noStyle name="autoLogin">
              自动登录
            </ProFormCheckbox>
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Login;
