using System;
using System.Linq;
using System.Text;
using System.Windows;
using System.Net.Sockets;
using System.Net;
using System.Threading;
using System.Windows.Threading;

namespace WpfApp1
{
    /// <summary>
    /// MainWindow.xaml 的交互逻辑
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();         
        }
        Socket socket;
        public void ButtonOne(object sender, RoutedEventArgs e)
        {
            try {
                socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
                IPAddress ipaddress = IPAddress.Parse(textBox.Text);
                int point=Convert.ToInt32(textBox1.Text);
                //EndPoint ipPoint = new IPEndPoint(ipaddress,point);
                IAsyncResult connResult =socket.BeginConnect(ipaddress,point, null, null);
                //尝试 链接ip   持续 1秒
                connResult.AsyncWaitHandle.WaitOne(1000, true);
                //判断该IP是否可连
                if (!connResult.IsCompleted)
                {
                    // 关闭socket
                    socket.Close();
                    MessageBox.Show("链接超时，请检查ip是否可连??", "ERROR");
                }
                else
                {
                   // socket.Connect(ipPoint);
                    // 最后的 ！ 号一定不能少
                    ShowMsg("连接成功!");
                    //开启线程  不断的接受
                    Thread th = new Thread(Recive);
                    th.IsBackground = true;
                    th.Start();
                }         
            }
            catch
            {
                MessageBox.Show("连接失败,请检查ip和端口￣へ￣", "ERROR");
            }
        }
        public void ShowMsg(string str)
        {
            //指定委托,解决不能 跨线程访问 textbox
            Application.Current.Dispatcher.BeginInvoke(DispatcherPriority.Normal,new Action(() =>
                 {                   
                     textBox2.AppendText(str + "\n");
                     //textBox2 文本框的内容  自动滚动到末尾
                     textBox2.ScrollToEnd();
                 }));          
        }
        //接受函数
        public void Recive()
        {
            while (true)
            {
                try
                {
                    //给于 适当的延时 ,根据 对方 发送消息 的 时间间隔 自行设置
                    Delay(12);
                    byte[] buffer = new byte[1024 * 1024 * 3];
                    int r = socket.Receive(buffer);
                    //检测是否 接受到消息
                    if (r == 0)
                    {
                        break;
                    }
                    else
                    {
                        string s = Encoding.Default.GetString(buffer, 0, r);       
                        ShowMsg(socket.RemoteEndPoint + ":" + s);                                                        
                    }
                }
                catch { }
            }
        }
        //发送 AT 指令
        public void ButtonTwo(object sender, RoutedEventArgs e)
        {
            string str = textBox3.Text.Trim();
            byte[] buffer=Encoding.UTF8.GetBytes(str);
            socket.Send(buffer);
        }
        //清空文本框 ，文本框置空即可
        public void ButtonThree(object sender, RoutedEventArgs e)
        {
            //  Application.Current.Dispatcher.BeginInvoke(DispatcherPriority.Normal, new Action(() =>
            //    {            }));
            textBox2.Text = "";
        }
        /*
         * 保存数据文件
         * */
        public void ButtonFour(object sender, RoutedEventArgs e)
        {
               string s=textBox2.Text;
                //保存 目录 也可自行设置
               string Path = "./data.txt";
               if (!System.IO.File.Exists(Path))
               {
                //如果不存在改文件 就创建文件data.txt
                   System.IO.FileStream f = System.IO.File.Create(Path);
                   f.Close();
                   f.Dispose();
               }
               //用来 获取本地 时间
              DateTime.Now.ToShortTimeString();
              DateTime dt = DateTime.Now;
              
              System.IO.StreamWriter f2 = new System.IO.StreamWriter(Path, true, Encoding.Default);
              if (s.Contains('!'))
              {
                string[] sArray = s.Split('!');
                f2.WriteLine(dt.ToLocalTime().ToString()+"---保存的数据\r\n" +sArray[1]);
                f2.Close();
              }
              else
              {
               f2.WriteLine(dt.ToLocalTime().ToString() + "---保存的数据\r\n" +s);
               f2.Close();
              }
               //StreamWriter必须显示调用Dispose()，否则数据肯定会丢失。                         
               f2.Dispose();
               //保存 后清空  防止下次保存重复数据
               textBox2.Text = "";   
        }
        
        //毫秒延时
        public static void Delay(int milliSecond)
        {
            int start = Environment.TickCount;
            while (Math.Abs(Environment.TickCount - start) < milliSecond)
            {
                ;
            }
        }
    
    }
}